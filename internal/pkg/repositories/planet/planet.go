package planet

import (
	"api/internal/pkg/model/entities"
	"api/internal/pkg/model/response"
	"api/internal/pkg/repositories/config"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanetRepository struct {
	instance *config.ConfigRepository
}

func NewPlanetRepository() *PlanetRepository {
	return &PlanetRepository{
		instance: config.GetConfigRepository(),
	}
}

func (r *PlanetRepository) Get() (*response.PlanetListResponse, error) {

	//Struct para pegar o id do mongo
	type HexId struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	var hexId HexId
	var planet response.PlanetResponse
	planets := new(response.PlanetListResponse)

	cursor, err := r.instance.Db.Database("dev").Collection("planets").Find(nil, bson.D{})
	if err != nil {
		defer cursor.Close(nil)
		return planets, err
	}

	for cursor.Next(nil) {
		err := cursor.Decode(&planet)
		err = cursor.Decode(&hexId)
		if err != nil {
			return planets, err
		}
		planet.Id = hexId.ID.Hex()
		planets.Planets = append(planets.Planets, planet)
	}

	return planets, nil
}

func (r *PlanetRepository) GetByName(name string) (*response.PlanetListResponse, error) {
	//todo implement mongodb query
	type HexId struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	println("GetByName")
	//matchStage := bson.D{{"$match", bson.D{{"name", name}}}}
	matchStage := bson.M{"name": name}
	//lookupStage := bson.D{{"$lookup",
	//	bson.D{{"from", "planets"},
	//		{"localField", "name"},
	//		{"as", "planets"}}}}
	var hexId HexId
	var planet response.PlanetResponse
	planets := new(response.PlanetListResponse)

	cursor, err := r.instance.Db.Database("dev").Collection("planets").Find(nil,
		matchStage)
	if err != nil {
		return nil, err
	}
	for cursor.Next(nil) {
		err := cursor.Decode(&planet)
		err = cursor.Decode(&hexId)
		if err != nil {
			return planets, err
		}
		planet.Id = hexId.ID.Hex()
		planets.Planets = append(planets.Planets, planet)
	}

	//if err = showLoadedCursor.All(nil, &list.Planets); err != nil {
	//	return nil, err
	//}

	return planets, nil
}

func (r *PlanetRepository) GetById(strID string) (*response.PlanetListResponse, error) {
	println("GetById")
	id, err := primitive.ObjectIDFromHex(strID)
	if err != nil {
		return nil, err
	}

	query := bson.M{"_id": bson.M{"$eq": id}}
	showLoadedCursor, err := r.instance.Db.Database("dev").Collection("planets").Find(nil, query)

	if err != nil {
		return nil, err
	}

	list := new(response.PlanetListResponse)
	if err = showLoadedCursor.All(nil, &list.Planets); err != nil {
		return nil, err
	}
	return list, nil
}

func (r *PlanetRepository) Insert(planet *entities.Planet) error {

	db := r.instance.Db.Database("dev")
	planetCollection := db.Collection("planets")
	_, err := planetCollection.InsertOne(nil, planet)
	if err != nil {

		return err
	}

	return nil
}

func (r *PlanetRepository) Delete(strID string) (*response.PlanetResponse, error) {
	//todo implement mongodb delete
	id, err := primitive.ObjectIDFromHex(strID)
	if err != nil {
		return nil, err
	}

	println("func Delete")
	query := bson.M{"_id": id}
	fmt.Println(query)
	db := r.instance.Db.Database("dev")
	planetCollection := db.Collection("planets")
	res, err := planetCollection.DeleteOne(nil, query)
	if err != nil {
		return nil, err
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", res.DeletedCount)
	return nil, nil
}

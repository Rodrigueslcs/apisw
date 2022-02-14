package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type ConfigRepository struct {
	Db *mongo.Client
}

var (
	instance *ConfigRepository
	Ctx      = context.TODO()
)

func newConfig() *ConfigRepository {
	host := os.Getenv("DB_ADDR")
	clientOptions := options.Client().ApplyURI("mongodb://" + host + ":27017/")
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &ConfigRepository{Db: client}
}

//GetConfigRepository retorna a instancia da conexão com o banco
func GetConfigRepository() *ConfigRepository {
	if instance == nil {
		instance = newConfig()
	}

	return instance
}

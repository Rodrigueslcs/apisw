package services

import (
	"api/internal/pkg/model/entities"
	"api/internal/pkg/model/request"
	"api/internal/pkg/model/response"
	"api/internal/pkg/repositories/planet"
	"net/http"
)

type PlanetService struct {
}

func NewTasksService() *PlanetService {
	return &PlanetService{}
}

func (s *PlanetService) GetPlanets(q *request.PlanetRequestQuery) (*response.PlanetListResponse, *response.ErrorResponse) {

	var err error

	r := planet.NewPlanetRepository()
	//Name query
	if len(q.Name) != 0 {
		println("getByName")
		rsp, err := r.GetByName(q.Name)
		if err != nil {
			return nil, response.NewErrorResponse(http.StatusBadRequest, err.Error())
		}
		return rsp, nil
	}
	//Id query
	if len(q.Id) != 0 {
		rsp, err := r.GetById(q.Id)
		if err != nil {
			return nil, response.NewErrorResponse(http.StatusBadRequest, err.Error())
		}
		return rsp, nil
	}
	//Get all
	rsp, err := r.Get()
	if err != nil {
		return nil, response.NewErrorResponse(http.StatusBadRequest, err.Error())
	}
	return rsp, nil
}

func (s *PlanetService) NewPlanet(rqtPlanet *request.PlanetRequest) *response.ErrorResponse {

	r := planet.NewPlanetRepository()
	p := new(entities.Planet)
	p.ParseFromRequest(rqtPlanet)

	err := r.Insert(p)
	if err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err.Error())
	}
	return nil
}
func (s *PlanetService) DeletePlanet(q *request.PlanetRequestQuery) *response.ErrorResponse {
	r := planet.NewPlanetRepository()
	if len(q.Id) != 0 {
		_, err := r.Delete(q.Id)
		if err != nil {
			return response.NewErrorResponse(http.StatusBadRequest, err.Error())
		}
		return nil
	}
	return nil

}

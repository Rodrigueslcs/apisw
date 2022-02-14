package entities

import (
	"api/internal/pkg/model/request"
	"api/internal/pkg/swapi"
)

type Planet struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Climate     string `json:"climate"`
	Terrain     string `json:"terrain"`
	Apparitions int    `json:"apparitions"`
}

func (p *Planet) ParseFromRequest(r *request.PlanetRequest) {
	apparitions, _ := swapi.GetQtdeFilm(r.Name)

	p.Name = r.Name
	p.Climate = r.Climate
	p.Terrain = r.Terrain
	p.Apparitions = apparitions
}

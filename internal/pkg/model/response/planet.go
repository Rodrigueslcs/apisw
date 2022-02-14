package response

type PlanetResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Climate     string `json:"climate"`
	Terrain     string `json:"terrain"`
	Apparitions int    `json:"apparitions"`
}

type PlanetListResponse struct {
	Planets []PlanetResponse `json:"planets"`
}

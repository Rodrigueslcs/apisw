package request

type PlanetRequest struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}

type PlanetRequestQuery struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

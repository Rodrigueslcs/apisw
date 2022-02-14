package swapi

import (
	//"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	URL = "https://swapi.dev/api/planets/"
)

type responseAPI struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []*planetAPI `json:"results"`
}
type planetAPI struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	ResidentURLs   []string `json:"residents"`
	FilmURLs       []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}

func newResponse() *responseAPI {
	return &responseAPI{
		Count:    0,
		Next:     "",
		Previous: "",
		Results:  make([]*planetAPI, 0),
	}
}

var slicePlanets []*planetAPI

func init() {
	slicePlanets = make([]*planetAPI, 0)
}
func GetQtdeFilm(planetName string) (int, error) {

	_, qtde, err := getPlanetsSWAPI(URL, planetName, 0)
	if err != nil {
		return 0, err
	}
	return qtde, nil
}

func getPlanetsSWAPI(planetsURL, planetName string, qtde int) (string, int, error) {

	req, err := http.NewRequest(http.MethodGet, planetsURL, nil)
	if err != nil {
		return "", 0, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	// fmt.Println(bytes.NewBuffer(body))
	response := newResponse()
	if err := json.Unmarshal(body, &response); err != nil {
		return "", 0, err
	}

	// fmt.Println(response.Results)

	for _, v := range response.Results {
		// fmt.Println(v)
		if strings.EqualFold(v.Name, planetName) {
			qtde += len(v.FilmURLs)

		}
	}
	if len(response.Next) > 1 {
		return "", qtde, nil

	}
	return getPlanetsSWAPI(response.Next, planetName, qtde)
}

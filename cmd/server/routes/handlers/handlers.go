package handlers

import (
	"api/internal/pkg/model/request"
	"api/internal/pkg/services"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Handler struct {
	PlanetService *services.PlanetService
}

func NewHandlers() *Handler {
	return &Handler{}
}

func (h *Handler) NewPlanet(c *fiber.Ctx) error {

	var rqtPlanet *request.PlanetRequest
	if err := json.Unmarshal(c.Request().Body(), &rqtPlanet); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	if err := h.PlanetService.NewPlanet(rqtPlanet); err != nil {
		return c.Status(err.Status).JSON(err)
	}
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) GetPlanet(c *fiber.Ctx) error {

	//localhost:3456/v1/planet
	//localhost:3456/v1/planet?id=123123
	//localhost:3456/v1/planet?name=earth
	q := &request.PlanetRequestQuery{
		Id:   c.Query("id"),
		Name: c.Query("name"),
	}

	if res, err := h.PlanetService.GetPlanets(q); err != nil {
		return c.Status(err.Status).JSON(err)
	} else {
		return c.Status(http.StatusOK).JSON(res)
	}
}

func (h *Handler) DeletePlanet(c *fiber.Ctx) error {
	q := &request.PlanetRequestQuery{
		Id:   c.Query("id"),
		Name: c.Query("name"),
	}
	err := h.PlanetService.DeletePlanet(q)
	if err != nil {
		return c.Status(err.Status).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(nil)
}

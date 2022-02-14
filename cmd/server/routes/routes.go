package routes

import (
	"api/cmd/server/routes/handlers"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	App     *fiber.App
	Handler *handlers.Handler
}

func NewRoutes() *Routes {
	return &Routes{
		App:     fiber.New(),
		Handler: handlers.NewHandlers(),
	}
}

func (r *Routes) RegisterRoutesGet(group string) {
	api := r.App.Group(group)
	api.Get("planet", r.Handler.GetPlanet)
	api.Post("planet", r.Handler.NewPlanet)
	api.Delete("planet", r.Handler.DeletePlanet)
}

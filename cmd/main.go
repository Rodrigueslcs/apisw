package main

import (
	"api/cmd/server/routes"
)

func main() {
	runApplication()
}
func runApplication() {

	r := routes.NewRoutes()

	r.RegisterRoutesGet("v1")
	//port := os.Getenv("PORT")

	r.App.Listen(":3456")
}

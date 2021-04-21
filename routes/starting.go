package routes

import (
	"github.com/bardiesel/twitchhub-backend-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func StartingRoutes(app fiber.Router) {
	route := app.Group("/startings")
	startingController := &controllers.StartingController{}

	route.Get("/", startingController.Index)
	route.Post("/create", startingController.Create)
}

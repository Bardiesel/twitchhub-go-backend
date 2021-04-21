package routes

import (
	"github.com/bardiesel/twitchhub-backend-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func EmoteRoutes(app fiber.Router) {
	route := app.Group("/emotes")
	emotesController := &controllers.EmoteController{}

	route.Get("/", emotesController.Index)
	route.Post("/create", emotesController.Create)
}

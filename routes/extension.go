package routes

import (
	"github.com/bardiesel/twitchhub-backend-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func ExtensionRoutes(app fiber.Router) {
	route := app.Group("/extensions")
	extensionsController := &controllers.ExtensionController{}

	route.Get("/", extensionsController.Index)
	route.Post("/create", extensionsController.Create)
}

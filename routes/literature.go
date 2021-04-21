package routes

import (
	"github.com/bardiesel/twitchhub-backend-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func LiteratureRoutes(app fiber.Router) {
	route := app.Group("/literatures")
	literaturesController := &controllers.LiteratureController{}

	route.Get("/", literaturesController.Index)
	route.Post("/create", literaturesController.Create)
}

package routes

import (
	"github.com/bardiesel/twitchhub-backend-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func CommandRoutes(app fiber.Router) {
	route := app.Group("/commands")
	commandsController := &controllers.CommandController{}

	route.Get("/", commandsController.Index)
	route.Post("/create", commandsController.Create)
}

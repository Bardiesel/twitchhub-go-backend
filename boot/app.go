package boot

import (
	"github.com/bardiesel/twitchhub-backend-go/utils"

	"github.com/bardiesel/twitchhub-backend-go/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func BootApp() *fiber.App {

	err := godotenv.Load()
	utils.ErrorHandler(err, "env not imported")

	app := fiber.New()

	api := app.Group("/api", cors.New())
	routes.StartingRoutes(api)
	routes.LiteratureRoutes(api)
	routes.EmoteRoutes(api)
	routes.CommandRoutes(api)
	routes.ExtensionRoutes(api)
	return app

}

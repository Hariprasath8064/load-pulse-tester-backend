package routes

import (
	"github.com/gofiber/fiber/v2"
	"testerbackend/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/run-test", handlers.RunLoadTest)
}

package routes

import (
	"github.com/anwarminst/jwt-auth-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	// Test route to verify application setup
	app.Get("/", controllers.Hello)
	app.Post("/api/register", controllers.Register)
}

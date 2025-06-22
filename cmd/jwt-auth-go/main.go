package main

import (
	"fmt"

	"github.com/anwarminst/jwt-auth-go/database"
	"github.com/anwarminst/jwt-auth-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to the database
	db, err := database.ConnectionDB()
	if err != nil {
		panic("could not connect to db")
	}
	fmt.Println("Connection is successful db: ", db)

	// Initialize Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4200",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization",
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetUpRoutes(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("I'm a GET request!")
	// })

	// Start the server
	err = app.Listen(":8000")
	if err != nil {
		// If unable to start the server, panic
		panic("could not start server")
	} else {
		fmt.Println("App is running st 8000")
	}

}

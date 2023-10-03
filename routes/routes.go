package routes

import (
	"BACKEND-MICROSERVICE-AUTHENTICATION/controllerGraph"
	"BACKEND-MICROSERVICE-AUTHENTICATION/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//Home route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	//Token validation Route (middleware example)
	app.Get("/validate", middleware.Validate)

	app.All("/graphql", controllerGraph.GraphQLHandler)
}

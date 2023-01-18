package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello, %s 👋!", "World")
	return c.SendString(msg) // => Hello john 👋!
}

func SetupRoutes(app *fiber.App) {
	// Home
	app.Get("/", helloWorld)
	/*
		// Stories
		app.Get("/cities", controllers.CityStories)
		app.Get("/first_city", controllers.GetCityById)
	*/
}

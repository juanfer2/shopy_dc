package servers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	products_routes "github.com/juanfer2/shopy_dc_golang/src/products/infrastructure/routes"
	"github.com/juanfer2/shopy_dc_golang/src/shared/infrastructure/routes"
)

func StartServerApp() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:   "${pid} ${status} - ${method} ${path}\n",
		TimeZone: "America/Bogotá",
		Done: func(c *fiber.Ctx, logString []byte) {
			fmt.Sprintf("%s - %d", c.Request().RequestURI(), c.Response().StatusCode())
		},
	}))
	app.Use(cors.New())
	app.Static("/images", "./tmp/images")

	routes.SetupRoutes(app)
	products_routes.SetupRoutes(app)
	app.Listen(":4000")
}

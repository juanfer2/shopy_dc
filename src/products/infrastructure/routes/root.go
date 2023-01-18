package products_routes

import (
	"github.com/gofiber/fiber/v2"
	products_application "github.com/juanfer2/shopy_dc_golang/src/products/application"
	products_controller "github.com/juanfer2/shopy_dc_golang/src/products/infrastructure/controller"
	products_repositories "github.com/juanfer2/shopy_dc_golang/src/products/infrastructure/repositories"
)

func SetupRoutes(app *fiber.App) {
	productRepository := products_repositories.NewProductPostgresRepository()
	productService := products_application.NewProductUseCase(productRepository)
	productsController := products_controller.NewProductController(productService)

	app.Get("/products", productsController.All)
}

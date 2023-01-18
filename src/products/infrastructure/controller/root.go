package products_controller

import (
	"github.com/gofiber/fiber/v2"

	products_application "github.com/juanfer2/shopy_dc_golang/src/products/application"
	products_serializers "github.com/juanfer2/shopy_dc_golang/src/products/infrastructure/serializers"
)

type ProductsController struct {
	productsUseCase *products_application.ProductsUseCase
}

func NewProductController(productsUseCase *products_application.ProductsUseCase) *ProductsController {
	return &ProductsController{
		productsUseCase: productsUseCase,
	}
}

func (p *ProductsController) All(c *fiber.Ctx) error {
	products := p.productsUseCase.All()
	products_serializers.ProductsToJson(products)

	return c.JSON(products_serializers.ProductsToJson(products))
}

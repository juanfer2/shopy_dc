package products_serializers

import (
	"fmt"

	products_domain "github.com/juanfer2/shopy_dc_golang/src/products/domain"
	"github.com/juanfer2/shopy_dc_golang/src/shared/utilities"
)

type ProductsSerializer struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImageUrl    string `json:"imageUrl"`
}

func ProductToJson(product products_domain.Product) ProductsSerializer {
	productsSerializer := ProductsSerializer{Id: product.ID, Name: product.Name,
		Price: fmt.Sprintf("$%d", product.Price), ImageUrl: product.ImageUrl}

	return productsSerializer
}

func ProductsToJson(products []products_domain.Product) []ProductsSerializer {
	productsSerializer := []ProductsSerializer{}

	for _, product := range products {
		productsSerializer = append(productsSerializer, ProductsSerializer{Id: product.ID,
			Name: product.Name, Price: fmt.Sprintf("$%d", product.Price),
			ImageUrl: utilities.GetUrlServer() + "/images/products/" + product.ImageUrl})
	}

	return productsSerializer
}

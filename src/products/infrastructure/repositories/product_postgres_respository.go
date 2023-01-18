package products_repositories

import (
	products_domain "github.com/juanfer2/shopy_dc_golang/src/products/domain"
	"github.com/juanfer2/shopy_dc_golang/src/shared/infrastructure/persistence/postgres"
)

type ProductPostgresRepository struct {
	postgres.PostgresRepository[products_domain.Product]
}

func NewProductPostgresRepository() products_domain.ProductRepository {
	repository := ProductPostgresRepository{}
	repository.PostgresClient = postgres.CreateClientFactory()

	return &repository
}

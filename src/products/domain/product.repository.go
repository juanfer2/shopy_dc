package products_domain

type ProductRepository interface {
	FindBy(id int) Product
	Create(product Product) Product
	CreateInBatches(products []Product)
	All() []Product
	Delete(product Product) Product
	Where(query any, arg ...any) []Product
}

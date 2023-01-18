package products_application

import products_domain "github.com/juanfer2/shopy_dc_golang/src/products/domain"

type ProductsUseCase struct {
	productRepository products_domain.ProductRepository
}

func NewProductUseCase(productRepository products_domain.ProductRepository) *ProductsUseCase {
	return &ProductsUseCase{
		productRepository: productRepository,
	}
}

func (p *ProductsUseCase) Create(
	name string, description string, price int, imageUrl string,
) products_domain.Product {
	product := products_domain.NewProduct(name, description, price, imageUrl)
	return p.productRepository.Create(product)
}

func (p *ProductsUseCase) All() []products_domain.Product {
	return p.productRepository.All()
}

func (p *ProductsUseCase) DestroyById(id int) products_domain.Product {
	product := p.productRepository.FindBy(id)
	p.productRepository.Delete(product)
	return product
}

func (p *ProductsUseCase) FindById(id int) products_domain.Product {
	return p.productRepository.FindBy(id)
}

func (p *ProductsUseCase) CreateMultiplesProducts(products []products_domain.Product) {
	p.productRepository.CreateInBatches(products)
}

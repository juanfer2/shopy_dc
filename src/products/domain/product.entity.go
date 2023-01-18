package products_domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          int    `json:"id" gorm:"serializer:json"`
	Name        string `json:"name" gorm:"serializer:json"`
	Description string `json:"description" gorm:"serializer:json"`
	Price       int    `json:"price" gorm:"serializer:json"`
	ImageUrl    string `json:"imageUrl" gorm:"serializer:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewProduct(name string, description string, price int, imageUrl string) Product {
	return Product{
		Name:        name,
		Description: description,
		Price:       price,
		ImageUrl:    imageUrl,
	}
}

package persistence_domain

type Repository[T comparable] interface {
	FindBy(id int) T
	All() []T
}

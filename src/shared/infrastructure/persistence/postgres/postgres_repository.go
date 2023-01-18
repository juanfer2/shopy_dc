package postgres

import (
	persistence_domain "github.com/juanfer2/shopy_dc_golang/src/shared/infrastructure/persistence/domain"
)

type PostgresRepository[T comparable] struct {
	PostgresClient *PostgresClient
}

func NewPostgresRepository[T comparable](
	postgresClient *PostgresClient,
) persistence_domain.Repository[T] {
	return &PostgresRepository[T]{
		PostgresClient: postgresClient,
	}
}

func (pr *PostgresRepository[T]) Create(newRecord T) T {
	pr.PostgresClient.Conn().Create(&newRecord)

	return newRecord
}

func (pr *PostgresRepository[T]) FindBy(id int) T {
	var record T
	pr.PostgresClient.Conn().First(&record, id)

	return record
}

func (pr *PostgresRepository[T]) All() []T {
	var records []T
	pr.PostgresClient.Conn().Find(&records)

	return records
}

func (pr *PostgresRepository[T]) Delete(record T) T {
	pr.PostgresClient.Conn().Delete(&record)

	return record
}

func (pr *PostgresRepository[T]) Save(record T) T {
	pr.PostgresClient.Conn().Save(&record)

	return record
}

func (pr *PostgresRepository[T]) Update(record T, updateRecord T) T {
	pr.PostgresClient.Conn().Model(&record).Updates(updateRecord)

	return record
}

func (pr *PostgresRepository[T]) Where(query any, arg ...any) []T {
	var table []T
	pr.PostgresClient.Conn().Where(query, arg).Find(&table)

	return table
}

func (pr *PostgresRepository[T]) Not(query any, arg ...any) []T {
	var table []T
	pr.PostgresClient.Conn().Not(query, arg).Find(&table)

	return table
}

func (pr *PostgresRepository[T]) CreateInBatches(data []T) {
	pr.PostgresClient.Conn().CreateInBatches(data, 100)
}

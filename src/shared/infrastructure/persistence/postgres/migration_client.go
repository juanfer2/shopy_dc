package postgres

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/juanfer2/shopy_dc_golang/src/shared/utilities"
	_ "github.com/lib/pq"
)

type MigrationClient struct {
	Migrate *migrate.Migrate
}

func NewMigrationClient() MigrationClient {
	return MigrationClient{
		Migrate: getMigrate(),
	}
}

func getMigrate() *migrate.Migrate {
	configClient := CreateConfigClient()

	m, err := migrate.New(
		configClient.MigrationFolder,
		configClient.GetUrl(),
	)

	if err != nil {
		log.Fatal(err)
	}

	return m
}

func (migrationClient MigrationClient) GetPathFolder() string {
	root := utilities.GetRootFolder()

	return root + "/src/shared/infrastructure/persistence/postgres/migrations"
}

func (migrationClient MigrationClient) RunMigration() {
	err := migrationClient.Migrate.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

	if err != nil {
		log.Fatal(err)
	}
}

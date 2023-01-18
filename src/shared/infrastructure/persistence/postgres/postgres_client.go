package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	ConfigClient    *ConfigClient
	MigrationClient *MigrationClient
}

func NewPostgresClient(
	configClient *ConfigClient,
	migrationClient *MigrationClient,
) PostgresClient {
	return PostgresClient{
		ConfigClient:    configClient,
		MigrationClient: migrationClient,
	}
}

func (p *PostgresClient) Conn() *gorm.DB {
	DSN := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable port=%s host=%s",
		p.ConfigClient.Username,
		p.ConfigClient.Password,
		p.ConfigClient.Database,
		p.ConfigClient.Port,
		p.ConfigClient.Host,
	)
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  DSN,  // data source name, refer https://github.com/jackc/pgx
			PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		}), &gorm.Config{})

	if err != nil {
		log.Fatalf("💥 Error while reading config file %s", err)
	}

	log.Println("✨ Database connected")

	return db
}

func (p *PostgresClient) CreateDataBase() {
	url := fmt.Sprintf(
		"user=%s password=%s port=%s host=%s",
		p.ConfigClient.Username,
		p.ConfigClient.Password,
		p.ConfigClient.Port,
		p.ConfigClient.Host,
	)

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		query := fmt.Sprintf("CREATE DATABASE %s", p.ConfigClient.Database)

		_, err = db.Exec(query)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Successfully created database..")
		}
	}
}

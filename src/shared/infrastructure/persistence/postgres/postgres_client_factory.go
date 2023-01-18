package postgres

func CreateClientFactory() *PostgresClient {
	configClient := CreateConfigClient()
	migrationClient := NewMigrationClient()
	postgrestClient := NewPostgresClient(configClient, &migrationClient)

	return &postgrestClient
}

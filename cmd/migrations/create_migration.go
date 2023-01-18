package migrations

import (
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/juanfer2/shopy_dc_golang/src/shared/infrastructure/persistence/postgres"
	"github.com/spf13/cobra"
)

var CreateMigrationCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Get user events.",
	Run:     createMigration,
}

func createMigration(cmd *cobra.Command, args []string) {
	color.Blue(">>>Migrate start 🚀 <<<")
	migration := postgres.NewMigrationClient()
	pathName := migration.GetPathFolder()

	// migrate create -ext sql -dir src/shared/infrastructure/persistence/postgres/migrations -seq create_users_table
	_, err := exec.Command("migrate", "create", "-ext", "sql", "-dir", pathName,
		args[0]).Output()

	if err != nil {
		color.Red(err.Error())
		log.Fatal(err)
	}

	color.Green("  Migratation: create_post create successfuly 🎉")
	color.Yellow("  Path: " + migration.GetPathFolder())
	color.Blue(">>>Migrate end 🚀 <<<")
}

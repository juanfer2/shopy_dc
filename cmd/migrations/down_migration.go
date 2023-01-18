package migrations

import (
	"github.com/fatih/color"
	"github.com/juanfer2/shopy_dc_golang/src/shared/infrastructure/persistence/postgres"
	"github.com/spf13/cobra"
)

var DownMigrationCmd = &cobra.Command{
	Use:     "down",
	Aliases: []string{"d"},
	Short:   "Get user events.",
	Run:     downMigration,
}

func downMigration(cmd *cobra.Command, args []string) {
	color.Blue(">>>Migrate DOWN 🚀 <<<")
	m := postgres.NewMigrationClient()
	m.Migrate.Down()
	color.Green("  Migrations downs successfuly")
	color.Blue(">>>Migrate DOWN 🚀 <<<")
}

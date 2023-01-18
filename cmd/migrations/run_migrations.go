package migrations

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fatih/color"
	"github.com/juanfer2/shopy_dc_golang/src/shared/infrastructure/persistence/postgres"
	"github.com/spf13/cobra"
)

var RunMigrationCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r"},
	Short:   "Get user events.",
	Run:     runMigration,
}

func runMigration(cmd *cobra.Command, args []string) {
	color.Blue(">>>Migrate UP 🚀 <<<")

	m := postgres.NewMigrationClient()

	if len(args) > 0 {
		value := args[0]
		version, _ := strconv.Atoi(value)
		err := m.Migrate.Migrate(uint(version))

		if err != nil {
			color.Red(err.Error())
			log.Fatal(err)
		}
	} else {
		err := m.Migrate.Up()

		if err != nil {
			color.Red(err.Error())
			log.Fatal(err)
		}
	}

	color.Green("  Migrations runs successfuly")
	color.Blue(">>>Migrate UP 🚀 <<<")
}

func StringToUint(value string) uint {
	u64, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	wd := uint(u64)

	return wd
}

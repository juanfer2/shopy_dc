package cmd

import (
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test about this service",
	Long:  `Test about this service`,
	Run:   runTests,
}

func runTests(cmd *cobra.Command, args []string) {
	color.Blue(">>>Start Tests 🚀 <<<")

	_, err := exec.Command("go", "test", "./spec/...").Output()
	if err != nil {
		color.Red(err.Error())
		log.Fatal(err)
	}

	color.Green("  golang-clean/* was deleted")
	color.Blue(">>>Start Tests 🚀 <<<")
}

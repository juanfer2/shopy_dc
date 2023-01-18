package cmd

import (
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cleanBuildCmd = &cobra.Command{
	Use:   "clean:build",
	Short: "clean about this service",
	Long:  `clean about this service`,
	Run:   cleanBuild,
}

func cleanBuild(cmd *cobra.Command, args []string) {
	color.Blue(">>>Clean Build Project start 🚀 <<<")

	_, err := exec.Command("rm", "-r", "golang-clean").Output()
	if err != nil {
		color.Red(err.Error())
		log.Fatal(err)
	}
	color.Green("  golang-clean/* was deleted")
	color.Blue(">>>Clean Build Project start 🚀 <<<")
}

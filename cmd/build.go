package cmd

import (
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build project",
	Long:  `build project`,
	Run:   buildProject,
}

func buildProject(cmd *cobra.Command, args []string) {
	color.Blue(">>>Building Project start 🚀 <<<")

	_, err := exec.Command("go", "build", "-o", "dist/golang-clean").Output()
	if err != nil {
		color.Red(err.Error())
		log.Fatal(err)
	}
	color.Green("  golang-clean/* was created successfully")
	color.Blue(">>>Building Project start 🚀 <<<")
}

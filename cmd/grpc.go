package cmd

import (
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "grpc about this service",
	Long:  `grpc about this service`,
	Run:   runGrpc,
}

func runGrpc(cmd *cobra.Command, args []string) {
	// protoc --go_out=. \                                                                                                        ‹system›
	//         --go-grpc_out=. \
	//         --go_opt=paths=source_relative \
	//         --go-grpc_opt=paths=source_relative \
	//         src/**/*.proto
	color.Blue(">>>Generate gRPC Protobuffers start 🚀 <<<")

	_, err := exec.Command(
		"sudo",
		"protoc", "--go_out=.",
		"--go-grpc_out=. ",
		"--go_opt=paths=source_relative ",
		"--go-grpc_opt=paths=source_relative ",
		"../**/*.proto",
	).Output()
	if err != nil {
		color.Red(err.Error())
		log.Fatal(err)
	}
	color.Blue(">>>Generate gRPC Protobuffers end 🚀 <<<")
}

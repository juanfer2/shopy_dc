package cmd

import (
	"github.com/spf13/cobra"

	"github.com/juanfer2/shopy_dc_golang/src/shared/infrastructure/servers"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server about this service",
	Long:  `server about this service`,
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	servers.StartServerApp()
}

var serverGrpcCmd = &cobra.Command{
	Use:   "server_grpc",
	Short: "server about this service",
	Long:  `server about this service`,
	Run:   runServerGrpc,
}

func runServerGrpc(cmd *cobra.Command, args []string) {
	servers.StartServerGrpcApp()
}

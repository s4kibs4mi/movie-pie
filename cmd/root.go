package cmd

import (
	"fmt"
	"github.com/s4kibs4mi/movie-pie/config"
	"github.com/s4kibs4mi/movie-pie/log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// RootCmd is the root command of golang-restful-boilerplate service
	RootCmd = &cobra.Command{
		Use:   "golang-restful-boilerplate",
		Short: "A grpc/http service",
		Long:  `An gRPC/HTTP JSON API backend service`,
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(migrationCmd)
}

// Execute executes the root command
func Execute() {
	if err := config.LoadConfig(); err != nil {
		fmt.Println("Failed to read config : ", err)
		os.Exit(1)
	}
	log.SetupLog()

	//if err := app.ConnectDB(); err != nil {
	//	log.Log().Printf("Failed to connect to database : %v\n", err)
	//	os.Exit(-1)
	//}

	if err := RootCmd.Execute(); err != nil {
		log.Log().Println(err)
		os.Exit(1)
	}
}

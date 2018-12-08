package cmd

import (
	"fmt"
	"github.com/codersgarage/golang-restful-boilerplate/app"
	"github.com/codersgarage/golang-restful-boilerplate/config"
	"github.com/codersgarage/golang-restful-boilerplate/log"
	"github.com/codersgarage/golang-restful-boilerplate/mq"
	"github.com/codersgarage/golang-restful-boilerplate/worker"
	"os"

	"github.com/spf13/cobra"
)

var (
	// RootCmd is the root command of boot-kit service
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

	if err := app.ConnectDB(); err != nil {
		log.Log().Fatalf("Failed to connect to database : %v\n", err)
		os.Exit(-1)
	}
	if err := mq.ConnectMQ(); err != nil {
		log.Log().Fatalf("Failed to connect to queue server: %v\n", err)
	}
	worker.RunWorker()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

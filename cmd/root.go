package cmd

import (
	"fmt"
	"github.com/codersgarage/golang-restful-boilerplate/app"
	"github.com/codersgarage/golang-restful-boilerplate/config"
	"github.com/codersgarage/golang-restful-boilerplate/log"
	"github.com/codersgarage/golang-restful-boilerplate/queue"
	"github.com/codersgarage/golang-restful-boilerplate/worker"
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

	if err := app.ConnectDB(); err != nil {
		log.Log().Printf("Failed to connect to database : %v\n", err)
		os.Exit(-1)
	}
	if err := queue.ConnectQueueServer(); err != nil {
		log.Log().Printf("Failed to connect to queue server: %v\n", err)
		os.Exit(-1)
	}
	worker.RunWorker()

	if err := worker.RegisterTasks(); err != nil {
		log.Log().Printf("Failed to register tasks: %v\n", err)
		os.Exit(-1)

	}

	if err := RootCmd.Execute(); err != nil {
		log.Log().Println(err)
		os.Exit(1)
	}
}

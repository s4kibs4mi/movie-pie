package cmd

import (
	"github.com/s4kibs4mi/movie-pie/cmd/migration"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "migration migrates database schemas",
}

func init() {
	migrationCmd.AddCommand(migration.MigAutoCmd)
}

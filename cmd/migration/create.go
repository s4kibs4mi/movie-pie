package migration

import (
	"github.com/codersgarage/golang-restful-boilerplate/app"
	"github.com/codersgarage/golang-restful-boilerplate/log"
	"github.com/codersgarage/golang-restful-boilerplate/models"
	"github.com/spf13/cobra"
)

var MigCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create creates database tables if not exists",
	Run:   create,
}

func create(cmd *cobra.Command, args []string) {
	monkey := &models.Monkey{}
	if err := app.DB().CreateTable(monkey).Error; err != nil {
		log.Log().Fatalf("Failed to create table [%s] with error {%v}", monkey.TableName(), err)
	}
}

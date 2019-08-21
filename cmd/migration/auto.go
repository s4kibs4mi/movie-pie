package migration

import (
	"github.com/s4kibs4mi/movie-pie/app"
	"github.com/s4kibs4mi/movie-pie/log"
	"github.com/s4kibs4mi/movie-pie/models"
	"github.com/spf13/cobra"
)

var MigAutoCmd = &cobra.Command{
	Use:   "auto",
	Short: "auto alter database tables if required",
	Run:   auto,
}

func auto(cmd *cobra.Command, args []string) {
	db := app.DB()

	u := models.User{}
	if err := db.Table(u.TableName()).AutoMigrate(&u); err != nil {
		log.Log().Errorln(err)
	}

	ss := models.Session{}
	if err := db.Table(ss.TableName()).AutoMigrate(&ss); err != nil {
		log.Log().Errorln(err)
	}

	m := models.Movie{}
	if err := db.Table(m.TableName()).AutoMigrate(&m); err != nil {
		log.Log().Errorln(err)
	}

	fm := models.FavouriteMovie{}
	if err := db.Table(fm.TableName()).AutoMigrate(&fm); err != nil {
		log.Log().Errorln(err)
	}
}

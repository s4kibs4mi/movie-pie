package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/s4kibs4mi/movie-pie/config"
	"github.com/s4kibs4mi/movie-pie/log"
)

var instance *gorm.DB

func ConnectDB() error {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DB().Username, config.DB().Password,
			config.DB().Host, config.DB().Port, config.DB().Name))
	if err != nil {
		return err
	}
	db.LogMode(true)
	db.SetLogger(log.Log())

	instance = db
	return nil
}

func DB() *gorm.DB {
	return instance
}

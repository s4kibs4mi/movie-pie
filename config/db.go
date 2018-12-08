package config

import (
	"github.com/spf13/viper"
)

// Database holds the database configuration
type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Options  map[string][]string
}

var db Database

// DB returns the default database configuration
func DB() Database {
	return db
}

// LoadDB loads database configuration
func LoadDB() {
	mu.Lock()
	defer mu.Unlock()

	db = Database{
		Name:     viper.GetString("database.name"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Options:  viper.GetStringMapStringSlice("database.options"),
	}
}

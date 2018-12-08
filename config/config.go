package config

import (
	"errors"
	"fmt"
	"sync"

	"github.com/spf13/viper"
	// this package is necessary to read config from remote consul
	_ "github.com/spf13/viper/remote"
)

var mu sync.Mutex

// LoadConfig initiates of config load
func LoadConfig() error {
	viper.BindEnv("CONSUL_URL")
	viper.BindEnv("CONSUL_PATH")

	consulURL := viper.GetString("CONSUL_URL")
	consulPath := viper.GetString("CONSUL_PATH")
	if consulURL == "" {
		return errors.New("CONSUL_URL missing")
	}
	if consulPath == "" {
		return errors.New("CONSUL_PATH missing")
	}

	viper.AddRemoteProvider("consul", consulURL, consulPath)
	viper.SetConfigType("yml")

	if err := viper.ReadRemoteConfig(); err != nil {
		return errors.New(fmt.Sprintf("%s named \"%s\"", err.Error(), consulPath))
	}

	LoadApp()
	LoadDB()
	LoadSvc()
	LoadMQ()
	LoadRedis()

	return nil
}

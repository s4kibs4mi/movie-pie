package config

import (
	"github.com/spf13/viper"
)

// Services holds the Services configuration
type Services struct {
	Service map[string]string
}

var svc Services

// Svc returns the default database configuration
func Svc(s string) string {
	return svc.Service[s]
}

// LoadSvc loads services configuration
func LoadSvc() {
	mu.Lock()
	defer mu.Unlock()

	svcs := viper.GetStringMapString("services")
	svs := map[string]string{}
	for s, v := range svcs {
		svs[s] = v
	}

	svc = Services{
		Service: svcs,
	}
}

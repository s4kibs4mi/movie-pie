package config

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/spf13/viper"
)

// RedisCfg holds the redis configuration
type RedisCfg struct {
	Address  string
	Password string
	DB       int
	Prefix   string
}

// URI build the redis uri from the configuration
func (r *RedisCfg) URI() string {
	u := url.URL{
		Scheme: "redis",
		Host:   r.Address,
		Path:   strconv.Itoa(r.DB),
	}
	if r.Password != "" {
		u.User = url.User(r.Password)
	}
	return u.String()
}

var redis RedisCfg

// Redis returns the default Redis configuration
func Redis() RedisCfg {
	return redis
}

// LoadRedis loads Redis configuration
func LoadRedis() {
	mu.Lock()
	defer mu.Unlock()

	redis = RedisCfg{
		Address:  fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		Prefix:   viper.GetString("redis.prefix"),
	}
}

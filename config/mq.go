package config

import (
	"github.com/spf13/viper"
)

type AMQP struct {
	Exchange      string
	ExchangeType  string
	BindingKey    string
	PrefetchCount int
}

type Worker struct {
	Name  string
	Count int
}

// Database holds the database configuration
type Mq struct {
	Broker        string
	DefaultQueue  string
	ResultBackend string
	AMQP          AMQP
	Worker        Worker
}

var mq Mq

// DB returns the default database configuration
func MQ() Mq {
	return mq
}

// LoadDB loads database configuration
func LoadMQ() {
	mu.Lock()
	defer mu.Unlock()

	mq = Mq{
		Broker:        viper.GetString("mq.broker"),
		DefaultQueue:  viper.GetString("mq.default_queue"),
		ResultBackend: viper.GetString("mq.result_backend"),
		AMQP: AMQP{
			Exchange:      viper.GetString("mq.amqp.exchange"),
			ExchangeType:  viper.GetString("mq.amqp.exchange_type"),
			BindingKey:    viper.GetString("mq.amqp.binding_key"),
			PrefetchCount: viper.GetInt("mq.amqp.prefetch_count"),
		},
		Worker: Worker{
			Name:  viper.GetString("mq.worker.name"),
			Count: viper.GetInt("mq.worker.count"),
		},
	}
}

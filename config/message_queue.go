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
type MessageQueue struct {
	Broker        string
	DefaultQueue  string
	ResultBackend string
	AMQP          AMQP
	Worker        Worker
}

var mq MessageQueue

// DB returns the default database configuration
func MQ() MessageQueue {
	return mq
}

// LoadDB loads database configuration
func LoadMQ() {
	mu.Lock()
	defer mu.Unlock()

	mq = MessageQueue{
		Broker:        viper.GetString("message_queue.broker"),
		DefaultQueue:  viper.GetString("message_queue.default_queue"),
		ResultBackend: viper.GetString("message_queue.result_backend"),
		AMQP: AMQP{
			Exchange:      viper.GetString("message_queue.amqp.exchange"),
			ExchangeType:  viper.GetString("message_queue.amqp.exchange_type"),
			BindingKey:    viper.GetString("message_queue.amqp.binding_key"),
			PrefetchCount: viper.GetInt("message_queue.amqp.prefetch_count"),
		},
		Worker: Worker{
			Name:  viper.GetString("message_queue.worker.name"),
			Count: viper.GetInt("message_queue.worker.count"),
		},
	}
}

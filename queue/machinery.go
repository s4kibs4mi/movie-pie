package queue

import (
	"github.com/RichardKnop/machinery/v1"
	cfg "github.com/RichardKnop/machinery/v1/config"
	"github.com/codersgarage/golang-restful-boilerplate/config"
)

var server *machinery.Server

func ConnectQueueServer() error {
	conf := &cfg.Config{
		Broker:        config.MQ().Broker,
		ResultBackend: config.MQ().ResultBackend,
		DefaultQueue:  config.MQ().DefaultQueue,
		AMQP: &cfg.AMQPConfig{
			Exchange:      config.MQ().AMQP.Exchange,
			ExchangeType:  config.MQ().AMQP.ExchangeType,
			BindingKey:    config.MQ().AMQP.BindingKey,
			PrefetchCount: config.MQ().AMQP.PrefetchCount,
		},
	}

	s, err := machinery.NewServer(conf)
	if err != nil {
		return err
	}
	server = s
	return nil
}

func QueueServer() *machinery.Server {
	return server
}

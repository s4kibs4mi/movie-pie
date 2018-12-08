package worker

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/codersgarage/golang-restful-boilerplate/config"
	"github.com/codersgarage/golang-restful-boilerplate/log"
	"github.com/codersgarage/golang-restful-boilerplate/mq"
)

var worker *machinery.Worker

func RunWorker() {
	go func() {
		worker = mq.MQServer().NewWorker(config.MQ().Worker.Name, config.MQ().Worker.Count)
		err := worker.Launch()
		if err != nil {
			log.Log().Errorln("Failed to launch worker :", err)
		}
	}()
}

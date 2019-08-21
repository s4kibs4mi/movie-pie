package worker

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/s4kibs4mi/movie-pie/config"
	"github.com/s4kibs4mi/movie-pie/log"
	"github.com/s4kibs4mi/movie-pie/queue"
	"os"
)

var worker *machinery.Worker

func RunWorker() {
	go func() {
		worker = queue.QueueServer().NewWorker(config.MQ().Worker.Name, config.MQ().Worker.Count)
		err := worker.Launch()
		if err != nil {
			log.Log().Errorln("Failed to launch worker :", err)
			os.Exit(-1)
		}
	}()
}

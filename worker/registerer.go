package worker

import (
	"github.com/s4kibs4mi/movie-pie/queue"
	"github.com/s4kibs4mi/movie-pie/tasks"
)

func RegisterTasks() error {
	if err := queue.QueueServer().RegisterTask(tasks.SayMonkeyTaskName, tasks.SayMonkey); err != nil {
		return err
	}
	return nil
}

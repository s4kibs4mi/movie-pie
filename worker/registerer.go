package worker

import (
	"github.com/codersgarage/golang-restful-boilerplate/queue"
	"github.com/codersgarage/golang-restful-boilerplate/tasks"
)

func RegisterTasks() error {
	if err := queue.QueueServer().RegisterTask(tasks.SayMonkeyTaskName, tasks.SayMonkey); err != nil {
		return err
	}
	return nil
}

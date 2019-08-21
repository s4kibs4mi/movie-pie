package tasks

import (
	"github.com/s4kibs4mi/movie-pie/log"
)

const (
	SayMonkeyTaskName = "say_monkey"
)

func SayMonkey() error {
	log.Log().Println("Received say monkey task event...")
	return nil
}

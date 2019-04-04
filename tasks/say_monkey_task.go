package tasks

import (
	"github.com/codersgarage/golang-restful-boilerplate/log"
)

const (
	SayMonkeyTaskName = "say_monkey"
)

func SayMonkey() error {
	log.Log().Println("Received say monkey task event...")
	return nil
}

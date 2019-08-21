package log

import (
	"github.com/s4kibs4mi/movie-pie/log/hooks"
	"github.com/sirupsen/logrus"
	"os"
)

var defLogger *logrus.Logger

func SetupLog() {
	defLogger = logrus.New()
	defLogger.Out = os.Stdout
	defLogger.AddHook(hooks.NewHook())
}

func getLevels() {

}

func Log() *logrus.Logger {
	return defLogger
}

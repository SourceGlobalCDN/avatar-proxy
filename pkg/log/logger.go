package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var GlobalLogger *logrus.Logger
var Level = logrus.InfoLevel

func Log() *logrus.Logger {
	if GlobalLogger == nil {
		GlobalLogger = NewLogger()
	}

	return GlobalLogger
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(NewFormatter())
	logger.SetOutput(os.Stdout)
	logger.SetLevel(Level)

	return logger
}

func SetLevel(level logrus.Level) {
	Level = level
	GlobalLogger = nil
	Log()
}

package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	Logger.SetOutput(os.Stdout)

	// log level
	Logger.SetLevel(logrus.InfoLevel)
}

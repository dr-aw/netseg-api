package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var levelColors = map[logrus.Level]string{
	logrus.DebugLevel: "\033[36m", // Blue
	logrus.InfoLevel:  "\033[32m", // Green
	logrus.WarnLevel:  "\033[33m", // Yellow
	logrus.ErrorLevel: "\033[31m", // Red
	logrus.FatalLevel: "\033[35m", // Purple
	logrus.PanicLevel: "\033[41m", // Red background
}

type CustomFormatter struct {
	logrus.TextFormatter
}

func InitLogger() {
	Logger.SetFormatter(&CustomFormatter{
		TextFormatter: logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
			// ForceColors:     true,
		},
	})

	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.InfoLevel)
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Take a layer
	layer, ok := entry.Data["layer"].(string)
	if !ok {
		layer = "UNKNOWN"
	}

	// Set colors
	color, exists := levelColors[entry.Level]
	if !exists {
		color = "\033[0m" // Сброс цвета
	}
	level := strings.ToUpper(entry.Level.String())
	layer = strings.ToUpper(layer)
	// Custom message
	logMessage := fmt.Sprintf(
		"%s%s | %s \033[1m[%s]\033[22m\t|\t%s\033[0m\n",
		color,
		entry.Time.Format(time.RFC3339),
		level,
		layer,
		entry.Message,
	)

	return []byte(logMessage), nil
}

// Logger global
var Logger = logrus.New()

// Adding a layer to the log entry
func LogWithLayer(layer string) *logrus.Entry {
	return Logger.WithField("layer", layer)
}

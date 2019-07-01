package log

import (
	"os"

	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	"github.com/sirupsen/logrus"
)

// New Create a new logger instance
func New(lvl logrus.Level) logrus.FieldLogger {
	return &logrus.Logger{
		Out: os.Stderr,
		Formatter: &prefixed.TextFormatter{
			QuoteEmptyFields: true,
			TimestampFormat:  "15:04:05",
			FullTimestamp:    true,
		},
		Hooks: make(logrus.LevelHooks),
		Level: lvl,
	}
}

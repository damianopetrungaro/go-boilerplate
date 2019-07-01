package log

import (
	"os"

	"github.com/damianopetrungaro/go-boilerplate/pkg/log"
	"github.com/sirupsen/logrus"
)

var lvl = os.Getenv("LOG_LEVEL")

// New Create a logger using the environment variable to define the log level
// it returns an error when the environment variable is an invalid log level
func New() (logrus.FieldLogger, error) {
	lvl, err := logrus.ParseLevel(lvl)
	if err != nil {
		return nil, err
	}
	return log.New(lvl), nil
}

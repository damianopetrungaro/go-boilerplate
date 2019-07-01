package log

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	tests := []struct {
		scenario string
		fun      func(t *testing.T)
	}{
		{
			scenario: "test logger is created with the log level from the env variable",
			fun:      testLevelFromEnvIsSet,
		},
		{
			scenario: "test logger is not created with an invalid log level",
			fun:      testLoggerIsNotCreatedWithAnInvalidLogLevel,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.fun(t)
		})
	}
}

func testLevelFromEnvIsSet(t *testing.T) {
	lvl = logrus.DebugLevel.String()
	log, err := New()
	assert.NoError(t, err)
	assert.Equal(t, logrus.DebugLevel, log.(*logrus.Logger).Level)
	lvl = ""
}

func testLoggerIsNotCreatedWithAnInvalidLogLevel(t *testing.T) {
	lvl = ""
	log, err := New()
	assert.EqualError(t, err, fmt.Sprintf("not a valid logrus Level: %q", lvl))
	assert.Equal(t, nil, log)
}

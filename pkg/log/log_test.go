package log

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
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
			scenario: "test logger is created with...",
			fun:      testLogger,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.fun(t)
		})
	}
}

func testLogger(t *testing.T) {
	log := New(logrus.DebugLevel).(*logrus.Logger)
	assert.Equal(t, logrus.DebugLevel, log.Level)
	assert.Equal(t, os.Stderr, log.Out)

	rw := bytes.NewBufferString("")
	log.SetOutput(rw)

	log.WithError(errors.New("an error occurred")).WithField("key", "value").WithField("key2", "").Debug("hey!")
	msg, err := ioutil.ReadAll(rw)
	assert.NoError(t, err)
	assert.Regexp(t,
		`^time="\d{2}:\d{2}:\d{2}" level=debug msg="hey!" error="an error occurred" key=value key2=""\n$`,
		string(msg),
	)
}

package server

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	tests := []struct {
		scenario string
		fun      func(t *testing.T)
	}{
		{
			scenario: "test server",
			fun:      testServerWithEnv,
		},
		{
			scenario: "test non existing routes",
			fun:      testServerEmptyEnv,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.fun(t)
		})
	}
}

func testServerWithEnv(t *testing.T) {
	address = ":80"
	readTimeout = "5"
	writeTimeout = "4"
	idleTimeout = "3"

	h := &http.ServeMux{}
	s := New(h)

	assert.Equal(t, s.Addr, ":80")
	assert.Equal(t, s.Handler, h)
	assert.Equal(t, s.ReadTimeout, time.Second*5)
	assert.Equal(t, s.WriteTimeout, time.Second*4)
	assert.Equal(t, s.IdleTimeout, time.Second*3)
}

func testServerEmptyEnv(t *testing.T) {
	address = ":99"
	readTimeout = ""
	writeTimeout = ""
	idleTimeout = ""

	h := &http.ServeMux{}
	s := New(h)

	assert.Equal(t, s.Addr, ":99")
	assert.Equal(t, s.Handler, h)
	assert.Equal(t, s.ReadTimeout, 0*time.Second)
	assert.Equal(t, s.WriteTimeout, 0*time.Second)
	assert.Equal(t, s.IdleTimeout, 0*time.Second)
}

package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {

	tests := []struct {
		scenario string
		fun      func(t *testing.T)
	}{
		{
			scenario: "test tracing",
			fun:      testTracing,
		},
		{
			scenario: "test non existing routes",
			fun:      testNonExistingRoute,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.fun(t)
		})
	}
}

func testTracing(t *testing.T) {
	traceID := "a trace id"
	log, _ := test.NewNullLogger()
	r := NewRouter(log)
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
	})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(traceIDHeader, traceID)
	r.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
	assert.Equal(t, traceID, res.Header.Get(traceIDHeader))
}

func testNonExistingRoute(t *testing.T) {
	log, _ := test.NewNullLogger()
	r := NewRouter(log)
	r.Get("/new", func(writer http.ResponseWriter, request *http.Request) {
	})
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/", nil))

	res := rec.Result()
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.NotEmpty(t, res.Header.Get(traceIDHeader))
}

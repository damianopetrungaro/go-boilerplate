package server

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	address      = os.Getenv("SERVER_ADDRESS")
	readTimeout  = os.Getenv("SERVER_READ_TIMEOUT")
	writeTimeout = os.Getenv("SERVER_WRITE_TIMEOUT")
	idleTimeout  = os.Getenv("SERVER_IDLE_TIMEOUT")
)

// New Return a new server
func New(h http.Handler) *http.Server {
	s := &http.Server{Addr: address, Handler: h}

	if i, err := strconv.Atoi(readTimeout); err == nil {
		s.ReadTimeout = time.Duration(i) * time.Second
	}

	if i, err := strconv.Atoi(writeTimeout); err == nil {
		s.WriteTimeout = time.Duration(i) * time.Second
	}

	if i, err := strconv.Atoi(idleTimeout); err == nil {
		s.IdleTimeout = time.Duration(i) * time.Second
	}

	return s
}

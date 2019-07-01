package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	logger "github.com/damianopetrungaro/go-boilerplate/internal/log"
)

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	log, err := logger.New()
	if err != nil {
		panic(err)
	}

	go func() {
		<-done
		defer cancel()
		signal.Stop(done)
		close(done)
		log.Info("Stopping the application...")
	}()

	r := NewApp(ctx, log)
	if err := r.Execute(); err != nil {
		log.WithError(err).Fatal("Could not gracefully shutdown the application.")
	}
	log.Info("Application stopped.")
}

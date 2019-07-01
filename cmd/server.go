package cmd

import (
	"context"
	"net/http"
	"time"

	database "github.com/damianopetrungaro/go-boilerplate/internal/db"
	"github.com/damianopetrungaro/go-boilerplate/internal/probe"
	"github.com/damianopetrungaro/go-boilerplate/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewServer Run a new server
func NewServer(ctx context.Context, log logrus.FieldLogger) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Run a server",
		RunE: func(cmd *cobra.Command, args []string) error {

			db, err := database.New()
			if err != nil {
				log.WithError(err).Fatal("Could not connect to the database.")
				return err
			}

			r := server.NewRouter(log)
			r.Route("/", probe.NewRouter(db))
			srv := server.New(r)

			done := make(chan struct{}, 1)
			go func(done chan<- struct{}) {
				<-ctx.Done()

				log.Info("Stopping the server...")

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				srv.SetKeepAlivesEnabled(false)
				if err := srv.Shutdown(ctx); err != nil {
					log.WithError(err).Fatal("Could not gracefully shutdown the server.")
				}

				// If you have any metrics or logs that need to be read before the shut down, remove the comment to the next 3 lines
				// log.Info("Waiting metrics and log to be read.")
				// <-ctx.Done()
				// log.Info("Metrics and log should be read.")

				log.Info("Server stopped.")
				close(done)
			}(done)

			log.Info("Server running...")
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}
			<-done
			return nil
		},
	}
}

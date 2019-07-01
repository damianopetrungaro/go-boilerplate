package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

// NewRoot Create a...
func NewRoot(ctx context.Context, log logrus.FieldLogger) *cobra.Command {
	cmd := &cobra.Command{
		Short: "app",
	}
	cmd.AddCommand(NewServer(ctx, log))
	return cmd
}

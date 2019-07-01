// +build wireinject

package main

import (
	"context"

	"github.com/damianopetrungaro/go-boilerplate/cmd"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewApp(ctx context.Context, log logrus.FieldLogger) *cobra.Command {
	wire.Build(cmd.NewRoot)
	return &cobra.Command{}
}

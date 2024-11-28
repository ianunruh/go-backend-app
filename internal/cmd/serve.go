package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ianunruh/go-backend-app/internal/cmd/options"
)

type ServeOptions struct {
	*options.Options

	ListenAddr string
}

func NewServeCmd(rootOpts *options.Options) *cobra.Command {
	opts := &ServeOptions{
		Options: rootOpts,
	}

	cmd := &cobra.Command{
		Use: "serve",

		Short: "Run the HTTP API server",

		RunE: func(cmd *cobra.Command, args []string) error {
			return runServe(cmd.Context(), opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.ListenAddr, "listen-addr", "l", "", "address to listen on")

	return cmd
}

func runServe(ctx context.Context, opts *ServeOptions) error {
	ct, err := opts.NewContainer()
	if err != nil {
		return fmt.Errorf("creating container: %w", err)
	}
	defer ct.Close()

	return ct.RunServer(ctx)
}

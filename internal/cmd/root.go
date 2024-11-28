package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ianunruh/go-backend-app/internal/cmd/options"
)

func NewRootCmd() *cobra.Command {
	opts := &options.Options{}

	cmd := &cobra.Command{
		Use: "app",
	}

	flags := cmd.PersistentFlags()
	flags.StringVar(&opts.ConfigPath, "config", "", "path to config")
	flags.StringVar(&opts.LogLevel, "log-level", "", "zap log level to use")

	cmd.AddCommand(NewServeCmd(opts))

	return cmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

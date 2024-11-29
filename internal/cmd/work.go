package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ianunruh/go-backend-app/internal/cmd/options"
)

type WorkOptions struct {
	*options.Options
}

func NewWorkCmd(rootOpts *options.Options) *cobra.Command {
	opts := &WorkOptions{
		Options: rootOpts,
	}

	cmd := &cobra.Command{
		Use:  "work",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runWork(opts)
		},
	}

	return cmd
}

func runWork(opts *WorkOptions) error {
	ct, err := opts.NewContainer()
	if err != nil {
		return err
	}
	defer ct.Close()

	mux, err := ct.NewAsynqServeMux()
	if err != nil {
		return err
	}

	srv := ct.NewAsynqServer()
	return srv.Run(mux)
}

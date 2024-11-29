package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ianunruh/go-backend-app/internal/cmd/options"
)

type ScheduleOptions struct {
	*options.Options
}

func NewScheduleCmd(rootOpts *options.Options) *cobra.Command {
	opts := &ScheduleOptions{
		Options: rootOpts,
	}

	cmd := &cobra.Command{
		Use:  "schedule",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSchedule(opts)
		},
	}

	return cmd
}

func runSchedule(opts *ScheduleOptions) error {
	ct, err := opts.NewContainer()
	if err != nil {
		return err
	}
	defer ct.Close()

	scheduler, err := ct.NewAsynqScheduler()
	if err != nil {
		return err
	}

	return scheduler.Run()
}

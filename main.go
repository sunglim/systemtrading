package main

import (
	"github.com/spf13/cobra"

	"github.com/sunglim/systemtrading/internal"
	"github.com/sunglim/systemtrading/internal/metrics"
	"github.com/sunglim/systemtrading/internal/options"
)

func init() {
	metrics.RegisterMetrics()
}

func main() {
	opts := options.NewOptions()
	cmd := options.InitCommands
	cmd.Run = func(cmd *cobra.Command, args []string) {
		internal.Wrapper(opts)
	}
	opts.AddFlags(cmd)

	if err := opts.Parse(); err != nil {
		panic(err)
	}

	if err := opts.Validate(); err != nil {
		panic(err)
	}
}

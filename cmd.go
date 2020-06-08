package main

import (
	"github.com/spf13/cobra"
)

func getRootCmd(handler func(in string, out string, fields []string, version bool) error) *cobra.Command {
	var in string
	var out string
	var fields []string
	var version bool

	cmd := &cobra.Command{
		Use:          "data-stripper",
		Short:        "tool to remove fields from csv or ndjson files",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler(in, out, fields, version)
		},
	}

	cmd.Flags().StringVar(&in, "in", "", "input file path")
	cmd.Flags().StringVar(&out, "out", "", "output file path")
	cmd.Flags().StringSliceVar(&fields, "field", nil, "field to remove")
	cmd.Flags().BoolVar(&version, "version", false, "show version")

	return cmd
}

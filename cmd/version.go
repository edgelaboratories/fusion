package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"

	versionCmd = &cobra.Command{ //nolint:exhaustruct_v5
		Use:   "version",
		Short: "Print the version number",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

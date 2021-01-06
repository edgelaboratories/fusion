package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "latest"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("fusion %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

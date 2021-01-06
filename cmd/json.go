package cmd

import (
	"fmt"
	"os"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Merge JSON files together",
	Run: func(cmd *cobra.Command, args []string) {
		config.AddDriver(json.Driver)
		config.WithOptions(func(opts *config.Options) {
			opts.DumpFormat = config.JSON
		})

		if err := config.LoadFiles(args...); err != nil {
			fmt.Println("failed to load the input files:", err)
			os.Exit(1)
		}

		if err := writeToFile(outputFilename); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
	jsonCmd.Flags().StringVarP(&outputFilename, "out", "o", "", "output file")
	if err := jsonCmd.MarkFlagRequired("out"); err != nil {
		panic(err)
	}
	if err := jsonCmd.MarkFlagFilename("out", "json"); err != nil {
		panic(err)
	}
}

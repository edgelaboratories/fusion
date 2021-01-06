package cmd

import (
	"fmt"
	"os"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
	"github.com/spf13/cobra"
)

var tomlCmd = &cobra.Command{
	Use:   "toml",
	Short: "Merge TOML files together",
	Run: func(cmd *cobra.Command, args []string) {
		config.AddDriver(toml.Driver)
		config.WithOptions(func(opts *config.Options) {
			opts.DumpFormat = config.Toml
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
	rootCmd.AddCommand(tomlCmd)
	tomlCmd.Flags().StringVarP(&outputFilename, "out", "o", "", "output file")
	if err := tomlCmd.MarkFlagRequired("out"); err != nil {
		panic(err)
	}
	if err := tomlCmd.MarkFlagFilename("out", "toml"); err != nil {
		panic(err)
	}
}

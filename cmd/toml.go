package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
	"github.com/spf13/cobra"
)

var tomlCmd = &cobra.Command{
	Use:   "toml",
	Short: "Merge TOML files together",
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			color.Green("operating on TOML files")
		}

		cfg := config.New("toml")
		cfg.AddDriver(toml.Driver)
		cfg.WithOptions(func(opts *config.Options) {
			opts.DumpFormat = config.Toml
		})

		if err := merge(cfg, outputFilename, args...); err != nil {
			color.Red(err.Error())
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

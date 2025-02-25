package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/spf13/cobra"
)

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "Merge YAML files together",
	Run: func(_ *cobra.Command, args []string) {
		if verbose {
			color.Green("operating on YAML files")
		}

		cfg := config.New("yaml")
		cfg.AddDriver(yaml.Driver)
		cfg.WithOptions(func(opts *config.Options) {
			opts.DumpFormat = config.Yaml
		})

		if err := merge(cfg, outputFilename, args...); err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(yamlCmd)
	yamlCmd.Flags().StringVarP(&outputFilename, "out", "o", "", "output file")
	if err := yamlCmd.MarkFlagRequired("out"); err != nil {
		panic(err)
	}
	if err := yamlCmd.MarkFlagFilename("out", "yaml", "yml"); err != nil {
		panic(err)
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/spf13/cobra"
)

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "Merge YAML files together",
	Run: func(cmd *cobra.Command, args []string) {
		config.AddDriver(yaml.Driver)
		config.WithOptions(func(opts *config.Options) {
			opts.DumpFormat = config.Yaml
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
	rootCmd.AddCommand(yamlCmd)
	yamlCmd.Flags().StringVarP(&outputFilename, "out", "o", "", "output file")
	if err := yamlCmd.MarkFlagRequired("out"); err != nil {
		panic(err)
	}
	if err := yamlCmd.MarkFlagFilename("out", "yaml", "yml"); err != nil {
		panic(err)
	}
}

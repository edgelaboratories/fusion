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
		cfg := config.New("json")
		cfg.AddDriver(json.Driver)
		cfg.WithOptions(func(opts *config.Options) {
			opts.DumpFormat = config.JSON
		})

		if err := merge(cfg, outputFilename, args...); err != nil {
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

package cmd

import (
	"fmt"
	"os"

	"github.com/gookit/config/v2"
)

var outputFilename string

func writeToFile(out string) error {
	f, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("failed to create the output file: %w", err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println("failed to close the output file:", err)
		}
	}()

	_, err = config.WriteTo(f)
	if err != nil {
		return fmt.Errorf("failed to write to the output file: %w", err)
	}

	return nil
}

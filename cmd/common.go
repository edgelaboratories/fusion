package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/gookit/config/v2"
)

var outputFilename string

func merge(cfg *config.Config, outputFilename string, sourceFiles ...string) error {
	if err := cfg.LoadFiles(sourceFiles...); err != nil {
		return fmt.Errorf("failed to load the input files: %w", err)
	}

	if err := writeToFile(cfg, outputFilename); err != nil {
		return err
	}

	return nil
}

func writeToFile(w io.WriterTo, outputFilename string) error {
	f, err := os.Create(outputFilename)
	if err != nil {
		return fmt.Errorf("failed to create the output file: %w", err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println("failed to close the output file:", err)
		}
	}()

	_, err = w.WriteTo(f)
	if err != nil {
		return fmt.Errorf("failed to write to the output file: %w", err)
	}

	return nil
}

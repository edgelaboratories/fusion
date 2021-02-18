package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/gookit/config/v2"
)

var outputFilename string

func merge(cfg *config.Config, outputFilename string, sourceFiles ...string) error {
	if verbose {
		color.Green("loading and merging %d files together: %v", len(sourceFiles), sourceFiles)
	}

	if err := cfg.LoadFiles(sourceFiles...); err != nil {
		return fmt.Errorf("failed to load the input files: %w", err)
	}

	return writeToFile(cfg, outputFilename)
}

func writeToFile(w io.WriterTo, outputFilename string) error {
	if verbose {
		color.Green("creating output file: %s", outputFilename)
	}

	f, err := os.Create(outputFilename)
	if err != nil {
		return fmt.Errorf("failed to create the output file: %w", err)
	}

	defer func() {
		if verbose {
			color.Green("closing output file: %s", outputFilename)
		}
		if err = f.Close(); err != nil {
			color.Red("failed to close the output file:", err)
		}
	}()

	if verbose {
		color.Green("writing result to file: %s", outputFilename)
	}
	_, err = w.WriteTo(f)
	if err != nil {
		return fmt.Errorf("failed to write to the output file: %w", err)
	}

	return nil
}

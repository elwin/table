package main

import (
	"log"
	"os"
	"strings"

	"github.com/elwin/table/pkg"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	format := pflag.StringP("format", "f", "csv", "Format, supported values: csv, json")
	input := pflag.StringP("input-file", "i", "", "Read input from file")
	pflag.Parse()

	var parser pkg.Parser
	switch strings.ToLower(*format) {
	case "csv":
		parser = pkg.CSVParser{}
	case "json":
		parser = pkg.JSONParser{}
	default:
		return errors.Errorf(`"%s" is not a supported parser`, *format)
	}

	in := os.Stdin
	if *input != "" {
		inputFile, err := os.Open(*input)
		if err != nil {
			return errors.Wrap(err, "failed to open file")
		}

		in = inputFile
	}

	err := pkg.Format(parser, in, os.Stdout)
	if err != nil {
		return err
	}

	return nil
}

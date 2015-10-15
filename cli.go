package main

import (
	"flag"
	"fmt"
	"github.com/mijime/merje/remarshal"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (this *CLI) Run(args []string) int {
	var (
		format  string
		out     string
		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(this.errStream)

	flags.StringVar(&format, "format", "json", "format")
	flags.StringVar(&format, "f", "json", "format")
	flags.StringVar(&out, "out", "", "output file path")
	flags.StringVar(&out, "o", "", "output file path")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(this.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	err := this.execute(format, out, flags.Args())

	if err != nil {
		log.Print(err)
		return ExitCodeError
	}

	return ExitCodeOK
}

func (this *CLI) execute(format string, out string, targets []string) (err error) {
	var (
		data, result interface{}
		iBuf, oBuf   []byte
		conv         remarshal.Converter
		writer       io.Writer
		option       remarshal.Option
	)

	// Input
	for _, target := range targets {
		if isFileExists(target) {
			iBuf, err = ioutil.ReadFile(target)

			if err != nil {
				return err
			}
		} else {
			iBuf = []byte(target)
		}

		// Find FileName
		option = remarshal.Option{FileName: target}
		conv, err = remarshal.Lookup(option)

		if err != nil {
			return err
		}

		if conv == nil {
			// Find Format
			option = remarshal.Option{target, option.Format}
			conv, err = remarshal.Lookup(option)

			if err != nil {
				return err
			}
		}

		data, err = conv.Unmarshal(iBuf)

		if err != nil {
			return err
		}

		result = remarshal.Merge(result, data)
	}

	option = remarshal.Option{out, format}
	conv, err = remarshal.Lookup(option)

	if err != nil {
		return err
	}

	oBuf, err = conv.Marshal(result)

	if err != nil {
		return err
	}

	if out == "" {
		writer = this.outStream
	} else {
		writer, err = os.Create(out)
	}

	_, err = writer.Write(oBuf)
	return err
}

func isFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

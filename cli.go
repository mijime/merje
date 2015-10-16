package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/mijime/merje/merge"
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

type Options struct {
	Format  string `short:"f" long:"format" description:"input/output format" default:"json"`
	Output  string `short:"o" long:"out" description:"output path"`
	Type    string `short:"t" long:"type" description:"merge type" default:"sum"`
	Version bool   `short:"v" long:"version" description:"print a version"`
}

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (this *CLI) Run(args []string) int {
	var options Options

	// Define option flag parse
	targets, err := flags.ParseArgs(&options, args)

	if err != nil {
		return ExitCodeError
	}

	// Show version
	if options.Version {
		fmt.Fprintf(this.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if e := this.execute(options, targets[1:]); e != nil {
		log.Print(e)
		return ExitCodeError
	}

	return ExitCodeOK
}

func (this *CLI) execute(options Options, targets []string) (err error) {
	var (
		data, result interface{}
		iBuf, oBuf   []byte
		conv         remarshal.Converter
		writer       io.Writer
		rOptions     remarshal.Options
	)

	operator, err := merge.Lookup(merge.Options{options.Type})

	if err != nil {
		return err
	}

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
		rOptions = remarshal.Options{FileName: target}
		conv, err = remarshal.Lookup(rOptions)

		if err != nil {
			return err
		}

		if conv == nil {
			// Find Format
			rOptions = remarshal.Options{target, rOptions.Format}
			conv, err = remarshal.Lookup(rOptions)

			if err != nil {
				return err
			}
		}

		data, err = conv.Unmarshal(iBuf)

		if err != nil {
			return err
		}

		result = operator.Merge(result, data)
	}

	rOptions = remarshal.Options{options.Output, options.Format}
	conv, err = remarshal.Lookup(rOptions)

	if err != nil {
		return err
	}

	oBuf, err = conv.Marshal(result)

	if err != nil {
		return err
	}

	if options.Output == "" {
		writer = this.outStream
	} else {
		writer, err = os.Create(options.Output)
	}

	_, err = writer.Write(oBuf)
	return err
}

func isFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/sync/errgroup"

	"github.com/mijime/merje/pkg/aggregate"
)

func createFile(output string) (*os.File, error) {
	if len(output) == 0 {
		return os.Stdout, nil
	}

	return os.Create(output)
}

func openFiles(fileChan chan<- *os.File, files []string) error {
	defer close(fileChan)

	if len(files) == 0 {
		fileChan <- os.Stdin

		return nil
	}

	for _, f := range files {
		if f == "-" {
			fileChan <- os.Stdin

			continue
		}

		fp, err := os.Open(f)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}

		fileChan <- fp
	}

	return nil
}

func aggregateFiles(agg aggregate.Aggregator, fileChan chan *os.File) interface{} {
	var curr interface{}

	for rfp := range fileChan {
		next, err := agg.Decode(rfp)
		rfp.Close()

		if err != nil {
			log.Println(fmt.Errorf("failed to decode: %w", err))
			continue
		}

		curr = agg.Merge(curr, next)
	}

	return curr
}

var version = "dev"

func main() {
	var (
		decFormat   string
		encFormat   string
		mergeType   string
		output      string
		showVersion bool
	)

	flag.StringVar(&decFormat, "decode", "", "json/yaml/toml")
	flag.StringVar(&encFormat, "encode", "", "json/yaml/toml/template")
	flag.StringVar(&mergeType, "merge", "or", "or/and/xor")
	flag.StringVar(&output, "out", "", "")
	flag.BoolVar(&showVersion, "version", false, version)
	flag.Parse()

	if showVersion {
		log.Println(version)
		return
	}

	agg, err := aggregate.New(decFormat, encFormat, mergeType)
	if err != nil {
		log.Panic(fmt.Errorf("failed to create aggregator: %w", err))
	}

	eg := errgroup.Group{}
	fileChan := make(chan *os.File)

	eg.Go(func() error {
		data := aggregateFiles(agg, fileChan)

		wfp, err := createFile(output)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}

		defer wfp.Close()

		return agg.Encode(wfp, data)
	})

	files := flag.Args()

	eg.Go(func() error {
		return openFiles(fileChan, files)
	})

	if err := eg.Wait(); err != nil {
		log.Panic(err)
	}
}

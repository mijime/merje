package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
)

var format = flag.String("f", "", "output format")
var output = flag.String("o", "", "output path")

func main() {
	flag.Parse()
	inputs := flag.Args()

	var result, dataBuf interface{}

	if len(inputs) <= 0 {
		dec := json.NewDecoder(os.Stdin)
		dec.Decode(&result)
	} else {
		for _, input := range inputs {
			var reader io.Reader
			if isFilePath(input) {
				var err error
				reader, err = os.Open(input)

				if err != nil {
					panic(err)
				}
			} else {
				reader = bytes.NewBufferString(input)
			}

			dec := json.NewDecoder(reader)
			dec.Decode(&dataBuf)

			result = mergeInterface(result, dataBuf)
		}
	}

	var writer io.Writer

	if *output == "" {
		writer = os.Stdout
	} else {
		var err error
		writer, err = os.Create(*output)
		if err != nil {
			panic(err)
		}
	}

	if *format == "" {
		enc := json.NewEncoder(writer)
		enc.Encode(result)
	} else {
		tmpl, err := buildTemplate(*format)
		if err != nil {
			panic(err)
		} else {
			tmpl.Execute(writer, result)
		}
	}
}

func mergeInterface(prev, curr interface{}) interface{} {
	switch prev.(type) {
	case map[string]interface{}:
		if curr == nil {
			return prev
		}

		return mergeHash(prev.(map[string]interface{}), curr.(map[string]interface{}))
	default:
		return curr
	}
}

func mergeHash(prev, curr map[string]interface{}) map[string]interface{} {
	for k, v := range curr {
		switch v.(type) {
		case map[string]interface{}:
			if prev[k] == nil {
				prev[k] = v

			} else {
				prev[k] = mergeInterface(prev[k], curr[k])
			}

		default:
			prev[k] = v
		}
	}

	return prev
}

func buildTemplate(format string) (*template.Template, error) {
	funcMap := template.FuncMap{
		"split":   strings.Split,
		"join":    strings.Join,
		"replace": strings.Replace,
		"base":    path.Base,
		"dir":     path.Dir,
	}

	if _, err := os.Stat(format); err != nil {
		return template.New("variable").Funcs(funcMap).Parse(format)

	} else {
		name := path.Base(format)
		return template.New(name).Funcs(funcMap).ParseFiles(format)
	}
}

func isFilePath(i string) bool {
	_, err := os.Stat(i)
	return err == nil
}

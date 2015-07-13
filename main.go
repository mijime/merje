package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/mijime/merje/remarshal"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

type Options struct {
	InputFormat  string `short:"i" long:"input-format" default:"ext" description:"Input format < json | yaml | toml | ext >"`
	OutputFormat string `short:"f" long:"output-format" default:"json" description:"Output format < json | yaml | toml | template path >"`
	OutputPath   string `short:"o" long:"output" description:"Output path"`
}

var opts Options

func main() {
	inputs, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(2)
	}

	var result interface{}
	var inputFormat remarshal.Format
	var inputBuffer []byte

	switch opts.InputFormat {
	case "yaml":
		inputFormat = remarshal.YAML
	case "toml":
		inputFormat = remarshal.TOML
	case "json":
		inputFormat = remarshal.JSON
	default:
		inputFormat = remarshal.Unknown
	}

	if len(inputs) <= 0 {
		inputBuffer, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			panic(err)
		}

		result, err = remarshal.Unmarshal(inputBuffer, inputFormat)

	} else {
		for _, input := range inputs {
			var err error
			var _inputFormat remarshal.Format

			if isFilePath(input) {
				inputBuffer, err = ioutil.ReadFile(input)

				if err != nil {
					panic(err)
				}

				switch path.Ext(input) {
				case ".yml", ".yaml":
					_inputFormat = remarshal.YAML
				case ".tml", ".toml":
					_inputFormat = remarshal.TOML
				case ".json":
					_inputFormat = remarshal.JSON
				default:
					_inputFormat = remarshal.Unknown
				}
			} else {
				inputBuffer = []byte(input)
			}

			if inputFormat == remarshal.Unknown {
				dataBuf, _ := remarshal.Unmarshal(inputBuffer, _inputFormat)
				result = mergeInterface(result, dataBuf)

			} else {
				dataBuf, _ := remarshal.Unmarshal(inputBuffer, inputFormat)
				result = mergeInterface(result, dataBuf)
			}
		}
	}

	var writer io.Writer
	var outputFormat string

	if opts.OutputPath == "" {
		writer = os.Stdout
	} else {
		var err error
		writer, err = os.Create(opts.OutputPath)
		if err != nil {
			panic(err)
		}
	}

	switch opts.OutputFormat {
	case "yaml":
		outputFormat = "{{yaml .}}"
	case "toml":
		outputFormat = "{{toml .}}"
	case "json":
		outputFormat = "{{json .}}"
	default:
		outputFormat = opts.OutputFormat
	}

	tmpl, err := buildTemplate(outputFormat)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(writer, result)
	if err != nil {
		panic(err)
	}
}

func mergeInterface(prev, curr interface{}) interface{} {
	if prev == nil {
		return curr
	}

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
		if prev[k] == nil {
			prev[k] = v
			continue
		}

		switch v.(type) {
		case map[string]interface{}:
			prev[k] = mergeInterface(prev[k], curr[k])

		default:
			prev[k] = v
		}
	}

	return prev
}

func buildTemplate(outputFormatString string) (*template.Template, error) {
	funcMap := template.FuncMap{
		"split":   strings.Split,
		"join":    strings.Join,
		"replace": strings.Replace,
		"base":    path.Base,
		"dir":     path.Dir,
		"json":    jsonMarshal,
		"yaml":    yamlMarshal,
		"toml":    tomlMarshal,
	}

	if _, err := os.Stat(outputFormatString); err != nil {
		return template.New("variable").Funcs(funcMap).Parse(outputFormatString)

	} else {
		name := path.Base(outputFormatString)
		return template.New(name).Funcs(funcMap).ParseFiles(outputFormatString)
	}
}

func jsonMarshal(data interface{}) (result string, err error) {
	b, err := remarshal.Marshal(data, remarshal.JSON)
	return string(b), err
}

func yamlMarshal(data interface{}) (result string, err error) {
	b, err := remarshal.Marshal(data, remarshal.YAML)
	return string(b), err
}

func tomlMarshal(data interface{}) (result string, err error) {
	b, err := remarshal.Marshal(data, remarshal.TOML)
	return string(b), err
}

func isFilePath(i string) bool {
	_, err := os.Stat(i)
	return err == nil
}

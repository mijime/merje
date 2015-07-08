package main

import (
	"flag"
	"github.com/mijime/merje/remarshal"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

var inputFormatString = flag.String("if", "ext", "input format")
var outputFormatString = flag.String("of", "json", "output format")
var outputPath = flag.String("o", "", "output path")

func main() {
	flag.Parse()
	inputs := flag.Args()

	var result interface{}
	var inputFormat remarshal.Format
	var inputBuffer []byte

	switch *inputFormatString {
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

	if *outputPath == "" {
		writer = os.Stdout
	} else {
		var err error
		writer, err = os.Create(*outputPath)
		if err != nil {
			panic(err)
		}
	}

	switch *outputFormatString {
	case "yaml":
		outputFormat = "{{yaml .}}"
	case "toml":
		outputFormat = "{{toml .}}"
	case "json":
		outputFormat = "{{json .}}"
	default:
		outputFormat = *outputFormatString
	}

	tmpl, err := buildTemplate(outputFormat)
	if err != nil {
		panic(err)

	} else {
		tmpl.Execute(writer, result)
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

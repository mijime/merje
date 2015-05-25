package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

var decode = flag.String("decode", "json", "input format")
var format = flag.String("format", "{{json .}}", "output format")
var output = flag.String("output", "", "output path")

func main() {
	flag.Parse()
	inputs := flag.Args()

	var result interface{}

	if len(inputs) <= 0 {
		b, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			panic(err)
		}

		result = typeDecode(b, *decode)

	} else {
		for _, input := range inputs {
			var b []byte
			var decoder string

			if isFilePath(input) {
				var err error
				b, err = ioutil.ReadFile(input)

				if err != nil {
					panic(err)
				}
				switch path.Ext(input) {
				case ".yml", ".yaml":
					decoder = "yaml"
				case ".tml", ".toml":
					decoder = "toml"
				case ".json":
					decoder = "json"
				default:
					decoder = *decode
				}
			} else {
				b = []byte(input)
				decoder = *decode
			}

			dataBuf := typeDecode(b, decoder)
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

	tmpl, err := buildTemplate(*format)
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

func typeDecode(b []byte, decoder string) interface{} {
	var dataBuf interface{}

	switch decoder {
	case "yaml":
		err := yaml.Unmarshal(b, &dataBuf)
		if err != nil {
			panic(err)
		}
	case "toml":
		err := toml.Unmarshal(b, &dataBuf)
		if err != nil {
			panic(err)
		}
	default:
		err := json.Unmarshal(b, &dataBuf)
		if err != nil {
			panic(err)
		}
	}

	return dataBuf
}

func jsonEncode(v interface{}) string {
	buf, err := json.Marshal(v)

	if err != nil {
		panic(err)
	}

	return string(buf)
}

func yamlEncode(v interface{}) string {
	buf, err := yaml.Marshal(v)

	if err != nil {
		panic(err)
	}

	return string(buf)
}

func tomlEncode(v interface{}) string {
	buf := new(bytes.Buffer)
	enc := toml.NewEncoder(buf)
	err := enc.Encode(v)

	if err != nil {
		panic(err)
	}

	return buf.String()
}

func buildTemplate(format string) (*template.Template, error) {
	funcMap := template.FuncMap{
		"split":   strings.Split,
		"join":    strings.Join,
		"replace": strings.Replace,
		"base":    path.Base,
		"dir":     path.Dir,
		"json":    jsonEncode,
		"yaml":    yamlEncode,
		"toml":    tomlEncode,
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

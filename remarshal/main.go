// remarshal, a utility to convert between serialization formats.
// Copyright (C) 2014 Danyil Bohdan
// License: MIT
package remarshal

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"strings"
)

type Format int

const (
	TOML Format = iota
	YAML
	JSON
	Unknown
)

func Unmarshal(input []byte, inputFormat Format) (data interface{}, err error) {
	return unmarshal(input, inputFormat)
}

func Marshal(data interface{}, outputFormat Format) (result []byte, err error) {
	return marshal(data, outputFormat, true)
}

// convertMapsToStringMaps recursively converts values of type
// map[interface{}]interface{} contained in item to map[string]interface{}. This
// is needed before the encoders for TOML and JSON can accept data returned by
// the YAML decoder.
func convertMapsToStringMaps(item interface{}) (res interface{}, err error) {
	switch item.(type) {
	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		for k, v := range item.(map[interface{}]interface{}) {
			res[k.(string)], err = convertMapsToStringMaps(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	case []interface{}:
		res := make([]interface{}, len(item.([]interface{})))
		for i, v := range item.([]interface{}) {
			res[i], err = convertMapsToStringMaps(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	default:
		return item, nil
	}
}

// convertNumbersToInt64 recursively walks the structures contained in item
// converting values of the type json.Number to int64 or, failing that, float64.
// This approach is meant to prevent encoders from putting numbers stored as
// json.Number in quotes or encoding large intergers in scientific notation.
func convertNumbersToInt64(item interface{}) (res interface{}, err error) {
	switch item.(type) {
	case map[string]interface{}:
		res := make(map[string]interface{})
		for k, v := range item.(map[string]interface{}) {
			res[k], err = convertNumbersToInt64(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	case []interface{}:
		res := make([]interface{}, len(item.([]interface{})))
		for i, v := range item.([]interface{}) {
			res[i], err = convertNumbersToInt64(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	case json.Number:
		n, err := item.(json.Number).Int64()
		if err != nil {
			f, err := item.(json.Number).Float64()
			if err != nil {
				// Can't convert to Int64.
				return item, nil
			}
			return f, nil
		}
		return n, nil
	default:
		return item, nil
	}
}

func stringToFormat(s string) (f Format, err error) {
	switch strings.ToLower(s) {
	case "toml":
		return TOML, nil
	case "yaml":
		return YAML, nil
	case "json":
		return JSON, nil
	default:
		return Unknown, errors.New("cannot convert string to Format: '" +
			s + "'")
	}
}

// unmarshal decodes serialized data in the Format inputFormat into a structure
// of nested maps and slices.
func unmarshal(input []byte, inputFormat Format) (data interface{},
	err error) {
	switch inputFormat {
	case TOML:
		_, err = toml.Decode(string(input), &data)
	case YAML:
		err = yaml.Unmarshal(input, &data)
		if err == nil {
			data, err = convertMapsToStringMaps(data)
		}
	case JSON:
		decoder := json.NewDecoder(bytes.NewReader(input))
		decoder.UseNumber()
		err = decoder.Decode(&data)
		if err == nil {
			data, err = convertNumbersToInt64(data)
		}
	}
	if err != nil {
		return nil, err
	}
	return
}

// marshal encodes data stored in nested maps and slices in the Format
// outputFormat.
func marshal(data interface{}, outputFormat Format,
	indentJSON bool) (result []byte, err error) {
	switch outputFormat {
	case TOML:
		buf := new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(data)
		result = buf.Bytes()
	case YAML:
		result, err = yaml.Marshal(&data)
	case JSON:
		result, err = json.Marshal(&data)
		if err == nil && indentJSON {
			buf := new(bytes.Buffer)
			err = json.Indent(buf, result, "", "  ")
			result = buf.Bytes()
		}
	}
	if err != nil {
		return nil, err
	}
	return
}

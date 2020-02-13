package aggregate

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mijime/merje/pkg/convert/json"
	"github.com/mijime/merje/pkg/convert/template"
	"github.com/mijime/merje/pkg/convert/toml"
	"github.com/mijime/merje/pkg/convert/yaml"
	"github.com/mijime/merje/pkg/merge/and"
	"github.com/mijime/merje/pkg/merge/or"
	"github.com/mijime/merje/pkg/merge/xor"
)

type (
	decodeFunc func(io.Reader) (interface{}, error)
	encodeFunc func(io.Writer, interface{}) error
	mergeFunc  func(interface{}, interface{}) interface{}
)

type Aggregator interface {
	Decode(*os.File) (interface{}, error)
	Encode(*os.File, interface{}) error
	Merge(interface{}, interface{}) interface{}
}

type BaseAggregator struct {
	DecodeFormat string
	EncodeFormat string
	MergeFunc    mergeFunc
}

func New(decFormat, encFormat, mergeType string) (Aggregator, error) {
	mrg, err := detectMerger(mergeType)
	if err != nil {
		return nil, fmt.Errorf("failed to detect merge type: %w", err)
	}

	return &BaseAggregator{
		DecodeFormat: decFormat,
		EncodeFormat: encFormat,
		MergeFunc:    mrg,
	}, nil
}

func (a BaseAggregator) Decode(fp *os.File) (interface{}, error) {
	dec, err := detectDecoder(a.DecodeFormat, fp)
	if err != nil {
		return nil, fmt.Errorf("failed to detect decoder: %w", err)
	}

	return dec(fp)
}

func (a BaseAggregator) Encode(fp *os.File, data interface{}) error {
	enc, err := detectEncoder(a.EncodeFormat, fp)
	if err != nil {
		return fmt.Errorf("failed to detect encoder: %w", err)
	}

	return enc(fp, data)
}

func (a BaseAggregator) Merge(curr, next interface{}) interface{} {
	return a.MergeFunc(curr, next)
}

func detectMerger(name string) (mergeFunc, error) {
	switch name {
	case "or":
		return or.Merge, nil
	case "and":
		return and.Merge, nil
	case "xor":
		return xor.Merge, nil
	default:
		return nil, fmt.Errorf("unsupport merge type: %s", name)
	}
}

func detectEncoderByFormat(format string) (encodeFunc, error) {
	_, err := os.Stat(format)
	if err == nil {
		tmpl, err := template.New(format)
		if err != nil {
			return nil, fmt.Errorf("failed to build template: %w", err)
		}

		return tmpl.Encode, nil
	}

	switch format {
	case "json":
		return json.Encode, nil
	case "yaml":
		return yaml.Encode, nil
	case "toml":
		return toml.Encode, nil
	default:
		return nil, fmt.Errorf("unsupport encode type: %s", format)
	}
}

func detectEncoderByFile(fp *os.File) (encodeFunc, error) {
	ext := filepath.Ext(fp.Name())

	switch ext {
	case "", ".json":
		return json.Encode, nil
	case ".yml", ".yaml":
		return yaml.Encode, nil
	case ".tml", ".toml":
		return toml.Encode, nil
	default:
		return nil, fmt.Errorf("unsupport encode extension: %s", ext)
	}
}

func detectEncoder(format string, fp *os.File) (encodeFunc, error) {
	if len(format) > 0 {
		return detectEncoderByFormat(format)
	}

	return detectEncoderByFile(fp)
}

func detectDecoderByFormat(format string) (decodeFunc, error) {
	switch format {
	case "json":
		return json.Decode, nil
	case "yaml":
		return yaml.Decode, nil
	case "toml":
		return toml.Decode, nil
	default:
		return nil, fmt.Errorf("unsupport decode type: %s", format)
	}
}

func detectDecoderByFile(fp *os.File) (decodeFunc, error) {
	ext := filepath.Ext(fp.Name())

	switch ext {
	case "", ".json":
		return json.Decode, nil
	case ".yml", ".yaml":
		return yaml.Decode, nil
	case ".tml", ".toml":
		return toml.Decode, nil
	default:
		return nil, fmt.Errorf("unsupport decode extension: %s", ext)
	}
}

func detectDecoder(format string, fp *os.File) (decodeFunc, error) {
	if len(format) > 0 {
		return detectDecoderByFormat(format)
	}

	return detectDecoderByFile(fp)
}

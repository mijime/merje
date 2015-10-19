package remarshal

import (
	"errors"
	"github.com/mijime/merje/adapter"
)

// Factory is
var Factory = adapter.New()

// Converter is
type Converter interface {
	Unmarshal(input []byte) (interface{}, error)
	Marshal(data interface{}) ([]byte, error)
}

// Options is
type Options struct {
	FileName, Format string
}

// Lookup is
func Lookup(options Options) (Converter, error) {
	adapter, err := Factory.Lookup(options)

	if err != nil {
		return nil, err
	}

	if adapter == nil {
		return nil, errors.New("Not support. FileName: " + options.FileName + ", Format: " + options.Format)
	}

	converter, ok := adapter.(Converter)

	if !ok {
		return nil, errors.New("Fail cast converter. FileName: " + options.FileName + " Format: " + options.Format)
	}

	return converter, nil
}

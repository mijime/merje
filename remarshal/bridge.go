package remarshal

import (
	"errors"
	"github.com/mijime/merje/adapter"
)

var Factory = adapter.New()

type Converter interface {
	Unmarshal(input []byte) (interface{}, error)
	Marshal(data interface{}) ([]byte, error)
}

type Option struct {
	FileName, Format string
}

func Lookup(option Option) (Converter, error) {
	adapter, err := Factory.Lookup(option)

	if err != nil {
		return nil, err
	}

	if adapter == nil {
		return nil, nil
	}

	converter, ok := adapter.(Converter)

	if !ok {
		return nil, errors.New("Can't convert error: FileName: " + option.FileName + " Format: " + option.Format)
	}

	return converter, nil
}

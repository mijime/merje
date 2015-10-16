package merge

import (
	"errors"
	"github.com/mijime/merje/adapter"
)

var Factory = adapter.New()

type Operator interface {
	Merge(curr, next interface{}) interface{}
}

type Options struct {
	Type string
}

func Lookup(options Options) (Operator, error) {
	adapter, err := Factory.Lookup(options)

	if err != nil {
		return nil, err
	}

	if adapter == nil {
		return nil, errors.New("Not support Type: " + options.Type)
	}

	operator, ok := adapter.(Operator)

	if !ok {
		return nil, errors.New("Fail cast Operator Type: " + options.Type)
	}

	return operator, nil
}

package merge

import (
	"errors"
	"github.com/mijime/merje/adapter"
)

// Factory is
var Factory = adapter.New()

// Operator is
type Operator interface {
	Merge(curr, next interface{}) interface{}
}

// Options is
type Options struct {
	Type string
}

// Lookup is
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

package adapter

import (
	"sync"
)

// Adapter is
type Adapter interface {
	Lookup(option interface{}) interface{}
}

// Factory is
type Factory struct {
	adapters map[string]Adapter
	sync.Mutex
}

// New is
func New() *Factory {
	adapters := make(map[string]Adapter, 0)
	return &Factory{adapters: adapters}
}

// Lookup is
func (f *Factory) Lookup(option interface{}) (interface{}, error) {
	for _, adapter := range f.adapters {
		result := adapter.Lookup(option)

		if result != nil {
			return result, nil
		}
	}

	return nil, nil
}

// Regist is
func (f *Factory) Regist(name string, adapter Adapter) {
	f.Lock()
	defer f.Unlock()

	f.adapters[name] = adapter
}

// Deregist is
func (f *Factory) Deregist(name string) {
	f.Lock()
	defer f.Unlock()

	delete(f.adapters, name)
}

package adapter

import (
	"sync"
)

type Adapter interface {
	Lookup(option interface{}) interface{}
}

type AdapterFactory struct {
	adapters map[string]Adapter
	sync.Mutex
}

func New() *AdapterFactory {
	adapters := make(map[string]Adapter, 0)
	return &AdapterFactory{adapters: adapters}
}

func (this *AdapterFactory) Lookup(option interface{}) (interface{}, error) {
	for _, adapter := range this.adapters {
		result := adapter.Lookup(option)

		if result != nil {
			return result, nil
		}
	}

	return nil, nil
}

func (this *AdapterFactory) Regist(name string, adapter Adapter) {
	this.Lock()
	defer this.Unlock()

	this.adapters[name] = adapter
}

func (this *AdapterFactory) Deregist(name string) {
	this.Lock()
	defer this.Unlock()

	delete(this.adapters, name)
}

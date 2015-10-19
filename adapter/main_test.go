package adapter

import (
	"testing"
)

type TestAdapter struct{}

func (ta *TestAdapter) Lookup(option interface{}) interface{} {
	return ta
}
func (ta *TestAdapter) UseFunc() bool {
	return true
}

func TestRun__main(t *testing.T) {
	factory := New()
	factory.Regist("test", &TestAdapter{})
	adapter, _ := factory.Lookup(nil)
	if !adapter.(*TestAdapter).UseFunc() {
		t.Error("can't use func")
	}
}

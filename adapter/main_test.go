package adapter

import (
	"testing"
)

type TestAdapter struct{}

func (this *TestAdapter) Lookup(option interface{}) interface{} {
	return this
}
func (this *TestAdapter) UseFunc() bool {
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

package factory

import "testing"

func TestGo(t *testing.T) {
	result := newGo()
	if result.getType() != goLang {
		t.Error("Wrong Language type")
	}
}

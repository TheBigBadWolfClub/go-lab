package factory

import "testing"

func TestRust(t *testing.T) {

	result := newRust()
	if result.getType() != rustLang {
		t.Error("Wrong Language type")
	}
}

package builder

import (
	"testing"
)

func TestNewPersonBuilder(t *testing.T) {

	pb := NewPersonBuilder()
	pb.Lives().
		At("The Square").
		WithPostalCode("1234121")

	pb.Works().
		As("Software Engineer").
		For("RustIncs").
		In("Round Plaza")

	person := pb.Build()

	if person.address != "The Square" {
		t.Errorf("fail to build personal stuff")
	}

	if person.address != "Round Plaza" {
		t.Errorf("fail to build work stuff")
	}
}

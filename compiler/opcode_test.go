package compiler

import (
	"testing"
)

func TestOpcode_String(t *testing.T) {
	expected := "ASSIGN"
	actual := ASSIGN.String()
	if actual != expected {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

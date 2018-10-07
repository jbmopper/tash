package tash

import (
	"testing"
)

func TestTash(t *testing.T) {
	input := -5
	expected := ""
	actual := Tash(input)

	if actual != expected {
		t.Errorf("Expected output of %v to be %v but got %v.", input, expected, actual)
	}
}

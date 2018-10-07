package tash

import (
	"testing"
)

func TestTash(t *testing.T) {
	input := 620
	expected := "a0"
	actual := Tash(input)

	if actual != expected {
		t.Errorf("Expected output of %v to be %v but got %v.", input, expected, actual)
	}
}

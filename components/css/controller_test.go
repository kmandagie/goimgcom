package css

import (
	"testing"
)

func TestRemoveSpace(t *testing.T) {
	data := "a b"
	result := removeSpace(data)
	if result != "ab" {
		t.Errorf("got: %s, want: %s.", result, "ab")
	}
}

func TestRemoveNewLine(t *testing.T) {
	data := "a b\n"
	result := removeNewLine(data)
	if result != "a b" {
		t.Errorf("got: %s, want: %s.", result, "a b")
	}
}

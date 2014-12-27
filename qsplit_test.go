package qsplit

import (
	"testing"
)

func TestQsplit(t *testing.T) {
	// the empty string should come back as a nil
	qs := Split([]byte(""))
	if qs != nil {
		t.Errorf("Empty string should be nil but got '%v'", string(qs[0]))
	}
	// a single word should come back as itself
	qs = Split([]byte("foo"))
	if len(qs) != 1 {
		t.Errorf("qs should be len 1 but is %v", len(qs))
	}
	if string(qs[0]) != "foo" {
		t.Errorf("'foo' should come back as 'foo' but got '%v'", string(qs[0]))
	}
}

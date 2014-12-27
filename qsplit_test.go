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
	
	// all whitespace should also be nil
	qs = Split([]byte("  \t       \t\t  "))
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
	
	// two words
	qs = Split([]byte("foo bar"))
	if len(qs) != 2 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	if string(qs[0]) != "foo" {
		t.Errorf("'foo' should come back as 'foo' but got '%v'", string(qs[0]))
	}
	if string(qs[1]) != "bar" {
		t.Errorf("'bar' should come back as 'bar' but got '%v'", string(qs[0]))
	}
	
	// two words with leading space and extra interspacing
	qs = Split([]byte("     foo \t  bar"))
	if len(qs) != 2 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	if string(qs[0]) != "foo" {
		t.Errorf("'foo' should come back as 'foo' but got '%v'", string(qs[0]))
	}
	if string(qs[1]) != "bar" {
		t.Errorf("'bar' should come back as 'bar' but got '%v'", string(qs[1]))
	}
	
	// begins with quote
	qs = Split([]byte(`"foo bar" baz`))
	if len(qs) != 2 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	if string(qs[0]) != "foo bar" {
		t.Errorf("should be 'foo bar' but got '%v'", string(qs[0]))
	}
	if string(qs[1]) != "baz" {
		t.Errorf("should be 'baz' but got '%v'", string(qs[1]))
	}
	
	// ends with quote
	qs = Split([]byte(`foo 'bar baz'`))
	if len(qs) != 2 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	if string(qs[0]) != "foo" {
		t.Errorf("should be 'foo' but got '%v'", string(qs[0]))
	}
	if string(qs[1]) != "bar baz" {
		t.Errorf("should be 'bar baz' but got '%v'", string(qs[1]))
	}
	
	// unterminated quote
	qs = Split([]byte(`foo 'bar baz"`))
	if len(qs) != 2 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	if string(qs[0]) != "foo" {
		t.Errorf("should be 'foo' but got '%v'", string(qs[0]))
	}
	if string(qs[1]) != `bar baz"` {
		t.Errorf("should be 'bar baz\"' but got '%v'", string(qs[1]))
	}
	
	// looks like a quote but isn't (not on word boundary)
	qs = Split([]byte(`foo bar'baz' quux`))
	if len(qs) != 3 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	if string(qs[0]) != "foo" {
		t.Errorf("should be 'foo' but got '%v'", string(qs[0]))
	}
	if string(qs[1]) != "bar'baz'" {
		t.Errorf("should be bar'baz' but got '%v'", string(qs[1]))
	}
	if string(qs[2]) != "quux" {
		t.Errorf("should be 'quux' but got '%v'", string(qs[2]))
	}
	
	// non-ASCII quotes
	qs = Split([]byte(`‹foo "bar"› xyz «「1 2 3』» abc 「this is a test」`))
	if len(qs) != 5 {
		t.Errorf("qs should be len 5 but is %v", len(qs))
	}
	for i, teststr := range []string{`foo "bar"`, "xyz", "「1 2 3』", "abc", "this is a test"} {
		if string(qs[i]) != teststr {
			t.Errorf("qs[%v] should be `%v` but is `%v`", i, teststr, string(qs[i]))
		}
	}
	
	// quotes end on closing quote, not on closing quote + word boundary
	qs = Split([]byte(`foo "bar"baz`))
	if len(qs) != 3 {
		t.Errorf("qs should be len 3 but is %v", len(qs))
	}
	for i, teststr := range []string{"foo", "bar", "baz"} {
		if string(qs[i]) != teststr {
			t.Errorf("qs[%v] should be `%v` but is `%v`", i, teststr, string(qs[i]))
		}
	}
	
	// some of everything, including trailing space
	qs = Split([]byte(`foo 'bar baz' ‹lorem     ipsum›     «a b c d e» x y  `))
	if len(qs) != 6 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	for i, teststr := range []string{"foo", "bar baz", "lorem     ipsum", "a b c d e", "x", "y"} {
		if string(qs[i]) != teststr {
			t.Errorf("qs[%v] should be `%v` but is `%v`", i, teststr, string(qs[i]))
		}
	}
}

func TestQsplitstring(t *testing.T) {
	qs := SplitString([]byte("foo bar"))
	if len(qs) != 2 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	if qs[0] != "foo" {
		t.Errorf("'foo' should come back as 'foo' but got '%v'", string(qs[0]))
	}
	if qs[1] != "bar" {
		t.Errorf("'bar' should come back as 'bar' but got '%v'", string(qs[0]))
	}
}

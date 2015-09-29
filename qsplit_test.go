package qsplit

import (
	"testing"
)

var testBytes [][]byte

func init() {
	testBytes = append(testBytes, []byte("foo"))
	testBytes = append(testBytes, []byte("foo bar"))
	testBytes = append(testBytes, []byte("foo 'bar baz' quux"))
	testBytes = append(testBytes, []byte(`data {"firstName": "John", "lastName": "Smith", "isAlive": true, "age": 25, "address": {"streetAddress": "21 2nd Street", "city": "New York", "state": "NY", "postalCode": "10021-3100"}, "phoneNumbers": [{"type": "home", "number": "212 555-1234"}, {"type": "office", "number": "646 555-4567"}], "children": [], "spouse": null}`))
}

func TestLocations(t *testing.T) {
	// the empty string should come back as a nil
	cp := Locations([]byte(""))
	if cp != nil {
		t.Errorf("Empty string should be nil but got '%v'", cp)
	}
	// all whitespace should also be nil
	cp = Locations([]byte("  \t       \t\t  "))
	if cp != nil {
		t.Errorf("Empty string should be nil but got '%v'", cp)
	}
	// a single word should come back as a single-element
	// slice. positions should be {0,3} for "foo"
	cp = Locations([]byte("foo"))
	if len(cp) != 1 {
		t.Errorf("cp should be len 1 but is %v", len(cp))
	}
	if cp[0][0] != 0 {
		t.Errorf("cp[0][0] of 'foo' should be 0 but is %v", cp[0][0])
	}
	if cp[0][1] != 3 {
		t.Errorf("cp[0][1] of 'foo' should be 3 but is %v", cp[0][1])
	}
	// a two words should come back as a two-element
	// slice. positions should be {0,3},{4,7} for "foo bar"
	cp = Locations([]byte("foo bar"))
	if len(cp) != 2 {
		t.Errorf("cp should be len 1 but is %v", len(cp))
		t.Errorf("cp is %v", cp)
	}
	if cp[0][0] != 0 {
		t.Errorf("cp[0][0] of 'foo bar' should be 0 but is %v", cp[0][0])
	}
	if cp[0][1] != 3 {
		t.Errorf("cp[0][1] of 'foo bar' should be 3 but is %v", cp[0][1])
	}
	if cp[1][0] != 4 {
		t.Errorf("cp[1][0] of 'foo bar' should be 4 but is %v", cp[1][0])
	}
	if cp[1][1] != 7 {
		t.Errorf("cp[1][1] of 'foo bar' should be 7 but is %v", cp[1][1])
	}
	// a single quoted word should come back as a single-element
	// slice. positions should be {1,4} for "'foo'"
	cp = Locations([]byte("'foo'"))
	if len(cp) != 1 {
		t.Errorf("cp should be len 1 but is %v", len(cp))
		t.Errorf("cp is %v", cp)
	}
	if cp[0][0] != 1 {
		t.Errorf("cp[0][0] of 'foo' should be 1 but is %v", cp[0][0])
	}
	if cp[0][1] != 4 {
		t.Errorf("cp[0][1] of 'foo' should be 4 but is %v", cp[0][1])
	}
	// a single quoted word should come back as a single-element
	// slice. positions should be {1,4} for `"foo"`
	cp = Locations([]byte(`"foo"`))
	if len(cp) != 1 {
		t.Errorf("cp should be len 1 but is %v", len(cp))
		t.Errorf("cp is %v", cp)
	}
	if cp[0][0] != 1 {
		t.Errorf("cp[0][0] of \"foo\" should be 1 but is %v", cp[0][0])
	}
	if cp[0][1] != 4 {
		t.Errorf("cp[0][1] of \"foo\" should be 4 but is %v", cp[0][1])
	}
	// a single quoted word should come back as a single-element
	// slice. positions should be {3,6} for "『foo』"
	cp = Locations([]byte("『foo』"))
	if len(cp) != 1 {
		t.Errorf("cp should be len 1 but is %v", len(cp))
		t.Errorf("cp is %v", cp)
	}
	if cp[0][0] != 3 {
		t.Errorf("cp[0][0] of 『foo』 should be 3 but is %v", cp[0][0])
	}
	if cp[0][1] != 6 {
		t.Errorf("cp[0][1] of 『foo』 should be 6 but is %v", cp[0][1])
	}
	// two quoted words should come back as a single-element
	// slice. positions should be {1,8} for "'foo bar'"
	cp = Locations([]byte("'foo bar'"))
	if len(cp) != 1 {
		t.Errorf("cp should be len 1 but is %v", len(cp))
		t.Errorf("cp is %v", cp)
	}
	if cp[0][0] != 1 {
		t.Errorf("cp[0][0] of 'foo' should be 1 but is %v", cp[0][0])
	}
	if cp[0][1] != 8 {
		t.Errorf("cp[0][1] of 'foo' should be 8 but is %v", cp[0][1])
	}
	// two words, one quoted and one not, should come back as a two-element
	// slice. positions should be {1,4},{6,9} for "'foo' bar"
	cp = Locations([]byte("'foo' bar"))
	if len(cp) != 2 {
		t.Errorf("cp should be len 2 but is %v", len(cp))
		t.Errorf("cp is %v", cp)
	}
	if cp[0][0] != 1 {
		t.Errorf("cp[0][0] of \"'foo' bar\" should be 1 but is %v", cp[0][0])
	}
	if cp[0][1] != 4 {
		t.Errorf("cp[0][1] of \"'foo' bar\" should be 4 but is %v", cp[0][1])
	}
	if cp[1][0] != 6 {
		t.Errorf("cp[1][0] of \"'foo' bar\" should be 6 but is %v", cp[1][0])
	}
	if cp[1][1] != 9 {
		t.Errorf("cp[1][1] of \"'foo' bar\" should be 9 but is %v", cp[1][1])
	}
}

func TestQsplit(t *testing.T) {
	// the empty string should come back as a nil
	qs := ToBytes([]byte(""))
	if qs != nil {
		t.Errorf("Empty string should be nil but got '%v'", string(qs[0]))
	}

	// all whitespace should also be nil
	qs = ToBytes([]byte("  \t       \t\t  "))
	if qs != nil {
		t.Errorf("Empty string should be nil but got '%v'", string(qs[0]))
	}

	// a single word should come back as itself
	qs = ToBytes([]byte("foo"))
	if len(qs) != 1 {
		t.Errorf("qs should be len 1 but is %v", len(qs))
	}
	if string(qs[0]) != "foo" {
		t.Errorf("'foo' should come back as 'foo' but got '%v'", string(qs[0]))
	}

	// two words
	qs = ToBytes([]byte("foo bar"))
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
	qs = ToBytes([]byte("     foo \t  bar"))
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
	qs = ToBytes([]byte(`"foo bar" baz`))
	if len(qs) != 2 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	} else {
		if string(qs[0]) != "foo bar" {
			t.Errorf("should be 'foo bar' but got '%v'", string(qs[0]))
		}
		if string(qs[1]) != "baz" {
			t.Errorf("should be 'baz' but got '%v'", string(qs[1]))
		}
	}

	// ends with quote
	qs = ToBytes([]byte(`foo 'bar baz'`))
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
	qs = ToBytes([]byte(`foo 'bar baz"`))
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
	qs = ToBytes([]byte(`foo bar'baz' quux`))
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
	qs = ToBytes([]byte(`‹foo "bar"› xyz «「1 2 3』» abc 「this is a test」`))
	if len(qs) != 5 {
		t.Errorf("qs should be len 5 but is %v", len(qs))
	}
	for i, teststr := range []string{`foo "bar"`, "xyz", "「1 2 3』", "abc", "this is a test"} {
		if string(qs[i]) != teststr {
			t.Errorf("qs[%v] should be `%v` but is `%v`", i, teststr, string(qs[i]))
		}
	}

	// quotes end on closing quote, not on closing quote + word boundary
	qs = ToBytes([]byte(`foo "bar"baz`))
	if len(qs) != 3 {
		t.Errorf("qs should be len 3 but is %v", len(qs))
	}
	for i, teststr := range []string{"foo", "bar", "baz"} {
		if string(qs[i]) != teststr {
			t.Errorf("qs[%v] should be `%v` but is `%v`", i, teststr, string(qs[i]))
		}
	}

	// some of everything, including trailing space
	qs = ToBytes([]byte(`foo 'bar baz' ‹lorem     ipsum›     «a b c d e» x y  `))
	if len(qs) != 6 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	for i, teststr := range []string{"foo", "bar baz", "lorem     ipsum", "a b c d e", "x", "y"} {
		if string(qs[i]) != teststr {
			t.Errorf("qs[%v] should be `%v` but is `%v`", i, teststr, string(qs[i]))
		}
	}
}

func TestQsplitString(t *testing.T) {
	qs := ToStrings([]byte("foo bar"))
	if len(qs) != 2 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	if qs[0] != "foo" {
		t.Errorf("'foo' should come back as 'foo' but got '%v'", string(qs[0]))
	}
	if qs[1] != "bar" {
		t.Errorf("'bar' should come back as 'bar' but got '%v'", string(qs[1]))
	}
}

func TestQsplitStringByte(t *testing.T) {
	cmd, chunks := ToStringBytes([]byte("foo bar baz quux"))
	if len(chunks) != 3 {
		t.Errorf("chunks should be len 3 but is %v", len(chunks))
	}
	if cmd != "foo" {
		t.Errorf("first chunk should have come back as 'foo' but got '%v'", cmd)
	}
	for i, tchunk := range []string{"bar", "baz", "quux"} {
		if string(chunks[i]) != tchunk {
			t.Errorf("chunk %v should come back as '%v' but got '%v'", i, tchunk, string(chunks[i]))
		}
	}
}

func TestOnce(t *testing.T) {
	qs := Once([]byte("foo bar baz"))
	if len(qs) != 2 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	if string(qs[0]) != "foo" {
		t.Errorf("'foo' should come back as 'foo' but got '%v'", string(qs[0]))
	}
	if string(qs[1]) != "bar baz" {
		t.Errorf("'bar' should come back as 'bar' but got '%v'", string(qs[1]))
	}
	qs = Once([]byte("'foo bar' baz"))
	if len(qs) != 2 {
		t.Errorf("qs should be len 2 but is %v", len(qs))
	}
	if string(qs[0]) != "foo bar" {
		t.Errorf("'foo' should come back as 'foo' but got '%v'", string(qs[0]))
	}
	if string(qs[1]) != "baz" {
		t.Errorf("'bar' should come back as 'bar' but got '%v'", string(qs[1]))
	}
}

func BenchmarkLocations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, byt := range testBytes {
			Locations(byt)
		}
	}
}

func BenchmarkOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, byt := range testBytes {
			Once(byt)
		}
	}
}

func BenchmarkToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, byt := range testBytes {
			ToBytes(byt)
		}
	}
}

func BenchmarkToStringBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, byt := range testBytes {
			ToStringBytes(byt)
		}
	}
}

func BenchmarkToStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, byt := range testBytes {
			ToStrings(byt)
		}
	}
}

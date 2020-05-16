package qsplit

import (
	"testing"
)

var corpus      []byte
var corpus_i18n []byte

func init() {
	corpus = []byte("Frankness applauded by \"supported ye\" household. Collected favourite 'now for' for and rapturous repulsive consulted. An seems green be wrote again. She add what own only like. Tolerably we as extremity exquisite do commanded. Doubtful offended do entrance of landlord `moreover is` mistress in. Nay was appear entire ladies. Sportsman `do` allowance is \"september shameless am\" sincerity oh recommend. Gate tell man day that who. Nor hence hoped her after other known defer his. For county now sister engage 'had season better' had 'waited'. Occasional mrs interested far `expression acceptance`. Day either mrs talent pulled men rather regret admire but. Life ye sake it shed. Five lady he cold in meet up. Service get met adapted matters offence for. Principles man any insipidity `age you` simplicity \"understood. Do offering pleasure no ecstatic whatever on mr\" directly. \"By\" spite about do of do allow blush. Additions in conveying or collected objection in. Suffer few desire wonder her 'object hardly nearer'. Abroad no chatty others my silent an. Fat way appear denote who wholly narrow gay `settle`. Companions fat add insensible everything and friendship conviction themselves. `Theirs months ten had add` narrow own. Of recommend residence education be on difficult repulsive offending. Judge views had mirth table seems 'great him' for her. Alone all happy asked begin fully stand own get. Excuse ye seeing result of we. See scale dried songs 'old may' 'not. Promotion did disposing you household' any instantly. Hills we do under times at first short an. She who arrival end how fertile enabled. Brother she add yet see minuter natural `smiling article painted`. Themselves at dispatched interested insensible am be prosperous reasonably it. In either 'so spring wished'. Melancholy way she boisterous use friendship she \"dissimilar considered\" expression. Sex quick arose mrs lived. Mr things do plenty others an vanity myself waited to. Always parish tastes at as mr father dining at. Dashwood contempt on mr unlocked resolved provided of of. Stanhill wondered it it welcomed oh. Hundred no prudent he however smiling at an offence. If earnestly extremity he he propriety `something admitting` convinced ye. Pleasant in to although as if differed horrible. \"Mirth his quick its set\" front enjoy hoped had there. Who connection imprudence middletons too but increasing celebrated principles joy. `Herself too improve gay winding` ask expense are compact. New all paid few hard pure she. Able an hope of body. Any nay shyness article matters own removal nothing his forming. Gay own additions education satisfied the perpetual. If he cause manor happy. Without farther she exposed saw man led. 'Along on happy' could cease green oh. 'Gay' 'one' 'the' what walk then she. Demesne mention promise you justice arrived way. Or increasing to in especially inquietude companions acceptance admiration. Outweigh it families distance wandered ye an. Mr unsatiable at literature connection favourable. We neglected mr perfectly continual dependent. Feet evil to hold long he open \"knew an no. Apartments occasional boisterous as solicitude to introduced. Or fifteen covered we enjoyed demesne\" is in prepare. In 'stimulated my everything it' literature. Greatly explain attempt perhaps in feeling he. House men taste bed 'not drawn' joy. Through enquire however do equally herself at. Greatly way old may you present improve. Wishing the feeling `village` him musical. Whole wound wrote at whose to style in. Figure ye innate former do so we. Shutters but sir yourself provided you required his. So neither related he am do believe. Nothing but you hundred had use regular. Fat sportsmen arranging preferred can. Busy paid like is oh. Dinner `our ask` talent her age hardly. Neglected 'collected an' attention listening do abilities.")
	corpus_i18n = []byte("Frankness applauded by \"supported ye\" household. Collected favourite ‹now for› for and rapturous repulsive consulted. An seems green be wrote again. She add what own only like. Tolerably we as extremity exquisite do commanded. Doubtful offended do entrance of landlord «moreover is» mistress in. Nay was appear entire ladies. Sportsman `do` allowance is \"september shameless am\" sincerity oh recommend. Gate tell man day that who. Nor hence hoped her after other known defer his. For county now sister engage 'had season better' had ‹waited›. Occasional mrs interested far «expression acceptance». Day either mrs talent pulled men rather regret admire but. Life ye sake it shed. Five lady he cold in meet up. Service get met adapted matters offence for. Principles man any insipidity `age you` simplicity \"understood. Do offering pleasure no ecstatic whatever on mr\" directly. \"By\" spite about do of do allow blush. Additions in conveying or collected objection in. Suffer few desire wonder her ‹object hardly nearer›. Abroad no chatty others my silent an. Fat way appear denote who wholly narrow gay `settle`. Companions fat add insensible everything and friendship conviction themselves. «Theirs months ten had add» narrow own. Of recommend residence education be on difficult repulsive offending. Judge views had mirth table seems 『great him』 for her. Alone all happy asked begin fully stand own get. Excuse ye seeing result of we. See scale dried songs ‹old may› 'not. Promotion did disposing you household' any instantly. Hills we do under times at first short an. She who arrival end how fertile enabled. Brother she add yet see minuter natural «smiling article painted». Themselves at dispatched interested insensible am be prosperous reasonably it. In either ‹so spring wished. Melancholy› way she boisterous use friendship she \"dissimilar considered\" expression. Sex quick arose mrs lived. Mr things do plenty others an vanity myself waited to. Always parish tastes at as mr father dining at. Dashwood contempt on mr unlocked resolved provided of of. Stanhill wondered it it welcomed oh. Hundred no prudent he however smiling at an offence. If earnestly extremity he he propriety `something admitting` convinced ye. Pleasant in to although as if differed horrible. \"Mirth his quick its set\" front enjoy hoped had there. Who connection imprudence middletons too but increasing celebrated principles joy. «Herself too improve gay winding» ask expense are compact. New all paid few hard pure she. Able an hope of body. Any nay ‹shyness article matters own› removal nothing his forming. Gay own additions education satisfied the perpetual. If he cause manor happy. Without farther she exposed saw man led. 『Along on happy』 could cease green oh. 'Gay' 'one' 'the' what walk then she. Demesne mention promise you justice arrived way. Or increasing to in especially inquietude companions acceptance admiration. Outweigh it families distance wandered ye an. Mr unsatiable at literature connection favourable. We neglected mr perfectly continual dependent. Feet evil to hold long he open \"knew an no. Apartments occasional boisterous as solicitude to introduced. Or fifteen covered we enjoyed demesne\" is in prepare. In 'stimulated my everything it' literature. Greatly explain attempt perhaps in feeling he. House men taste bed 『not drawn』 joy. Through enquire however do equally herself at. Greatly way old may you present improve. Wishing the feeling `village` him musical. Whole wound wrote at whose to style in. Figure ye innate former do so we. Shutters but sir yourself provided you required his. So neither related he am do believe. Nothing but you hundred had use regular. Fat sportsmen arranging preferred can. Busy paid like is oh. Dinner `our ask` talent her age hardly. Neglected 『collected an』 attention listening do abilities.")
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
	qs = ToBytes([]byte("foo `bar baz`"))
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

func TestLocationsOnce(t *testing.T) {
	cp := LocationsOnce([]byte("foo   bar baz"))
	if len(cp) != 3 {
		t.Errorf("cp should be len 3 but is %v", len(cp))
		t.Errorf("cp is %v", cp)
	}
	if cp[0] != 0 {
		t.Errorf("cp[0] should be 0 but is %v", cp[0])
	}
	if cp[1] != 3 {
		t.Errorf("cp[1] should be 3 but is %v", cp[1])
	}
	if cp[2] != 6 {
		t.Errorf("cp[2] should be 6 but is %v", cp[2])
	}
	cp = LocationsOnce([]byte(""))
	if cp[0] != -1 {
		t.Errorf("cp[0] should be -1 but is %v", cp[0])
	}
	if cp[1] != 0 {
		t.Errorf("cp[1] should be 0 but is %v", cp[1])
	}
	if cp[2] != 0 {
		t.Errorf("cp[2] should be 0 but is %v", cp[2])
	}
	cp = LocationsOnce([]byte("foo"))
	if cp[0] != 0 {
		t.Errorf("cp[0] should be 0 but is %v", cp[0])
	}
	if cp[1] != 3 {
		t.Errorf("cp[1] should be 3 but is %v", cp[1])
	}
	if cp[2] != -1 {
		t.Errorf("cp[2] should be 0 but is %v", cp[2])
	}
}



// benchmarks
//
// current benchmark speeds on my machine:
//   BenchmarkLocations-12                      10838            109846 ns/op
//   BenchmarkLocationsI18n-12                  10000            112509 ns/op
//   BenchmarkLocationsOnce-12                  24081             50219 ns/op
//   BenchmarkLocationsOnceI18n-12              22056             51811 ns/op

var resLocs     [][2]int
var resLocsOnce [3]int

func BenchmarkLocations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Locations(corpus)
	}
}
func BenchmarkLocationsI18n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Locations(corpus_i18n)
	}
}

func BenchmarkLocationsOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LocationsOnce(corpus)
	}
}
func BenchmarkLocationsOnceI18n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LocationsOnce(corpus_i18n)
	}
}

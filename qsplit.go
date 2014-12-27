package qsplit

import (
	"bytes"
	"regexp"
)

var (
	spaceRE *regexp.Regexp
	quotes  map[rune]rune
)

func init() {
	spaceRE = regexp.MustCompile(`\pZ`)
	quotes  = map[rune]rune{
		'\'':'\'',
		'"':'"',
		'‹':'›',
		'«':'»',
	}
}

func Split(b []byte) [][]byte {
	var ss [][]byte // slice of slice of bytes
	var s  string   // temprary string
	var i, j int = 0, 0
	// we need to operate at the runes level
	r := bytes.Runes(b)
	for i = 0; i < len(r); i++ {
		if spaceRE.MatchString(string(r[i])) {
			// are we looking at a space
			if j == 0 {
				continue // yes. but not in a word; toss it
			}
			// yes & we were in a word, which has now ended
			ss = append(ss, []byte(s)) // append s to ss
			s = ""                     // reset s
			j = 0                      // reset j
			continue
		}
		// not looking at a space; see if we're at an opening quote
		c, ok := quotes[r[i]]
		if j == 0 && ok {
			i++ // yes, so move up one rune
			for i < len(r) && r[i] != c {
				// and string runes together until we hit the end quote
				s = s + string(r[i])
				i++
			}
			ss = append(ss, []byte(s))
			s = s[:0]
			continue
		}
		// we're in a plain old word then. increment j and append the
		// rune to s
		j++
		s = s + string(r[i])
	}
	// append to ss if the end of r was inside a word
	if j != 0 {
		ss = append(ss, []byte(s))
	}
	return ss
}

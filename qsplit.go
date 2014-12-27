/*
Package qsplit (short for "quoted split") performs a smart,
Unicode-aware split-on-whitespace. It returns a slice of the
non-whitespace "chunks" contained in a byte slice. It treats text
within balanced quotes as a single chunk.

Whitespace, according to qsplit, is `[\pZ\t]` (Unicode separators plus
horizontal tab).

Qsplit is aware of several quote character pairs:

    ASCII single: ''
    ASCII double: ""
    Guillemets:   ‹›, «»
    Japanese:     「」,『』

These are the rules used to delineate chunks of quoted text:

    * Quotes begin only at a word boundary
    * Quotes extend to the first closing quotation mark (regardless of
      word boundaries)
    * Quotes do not nest

*/
package qsplit

// Copyright (c) 2014 Shawn Boyette <shawn@firepear.net>. All rights
// reserved.  Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"bytes"
	"regexp"
)

var (
	spaceRE *regexp.Regexp
	quotes  map[rune]rune
)

func init() {
	spaceRE = regexp.MustCompile(`[\pZ\t]`)
	quotes  = map[rune]rune{
		'\'':'\'', '"':'"',
		'‹':'›', '«':'»',
		'「':'」', '『':'』',
	}
}

// Split performs a smart split-on-whitespace.
func Split(b []byte) [][]byte {
	var sb [][]byte // slice of slice of bytes
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
			sb = append(sb, []byte(s)) // append s to ss
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
			sb = append(sb, []byte(s))
			s = ""
			continue
		}
		// we're in a plain old word then. increment j and append the
		// rune to s
		j++
		s = s + string(r[i])
	}
	// append to ss if the end of r was inside a word
	if j != 0 {
		sb = append(sb, []byte(s))
	}
	return sb
}

// SplitString is a convenience function which works like Split, but
// returns a slice of strings.
func SplitString(b []byte) []string {
	var ss []string
	bslices := Split(b)
	for _, bslice := range bslices {
		ss = append(ss, string(bslice))
	}
	return ss
}

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
    * Quotes extend to the first closing quotation mark which matches
      the opening quote (regardless of word boundaries)
    * Quotes do not nest

*/
package qsplit

// Copyright (c) 2014,2015 Shawn Boyette <shawn@firepear.net>. All
// rights reserved.  Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"regexp"
	"unicode/utf8"
)

var (
	Version = "2.1.0" // current version
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

// Locations finds where the input byteslice would be split, and
// returns the beginning and end points of all text chunks which would
// be returned by one of the To...() functions.
func Locations(b []byte) ([][2]int) {

	// then add qsplit.Once()

	var si [][2]int // slice of tuples of ints
	var inw, inq, ok bool // in-word and in-quote flags; map test var
	var rune, endq rune   // current rune; end-quote for current quote chunk
	var i, idx int        // first index of chunk; byte index of current rune

	// we need to operate at the runes level
	runes := bytes.Runes(b)
	for _, rune = range runes {
		switch {
		case inq:
			// in a quoted chunk, if we're looking at the ending
			// quote, unset inq and append a the tuple for this chunk
			// to the return list.
			if rune == endq {
				inq = false
				si = append(si, [2]int{i, idx})
			}
		case spaceRE.MatchString(string(rune)):
			// if looking at a space and inw is set, end the present
			// chunk and append a new tuple. else just move on.
			if inw {
				inw = false
				si = append(si, [2]int{i, idx})
			}
		case inw:
			// if in a regular word, do nothing
		default:
			if endq, ok = quotes[rune]; ok {
				// looking at a quote; set inq and i
				inq = true
				i = idx + utf8.RuneLen(rune)
			} else {
				// looking at the first rune in a word. set inw& i
				inw = true
				i = idx
			}
		}
		idx += utf8.RuneLen(rune)
	}
	// append the tuple for the last chunk if we were still in a word
	// or quote
	if inw || inq {
		si = append(si, [2]int{i, idx})
	}
	return si
}

// ToBytes performs a quoted split to a slice of byteslices.
func ToBytes(b []byte) [][]byte {
	var sb [][]byte // slice of slice of bytes
	cp := Locations(b) // get chunk positions
	for _, pos := range cp {
		sb = append(sb, b[pos[0]:pos[1]])
	}
	return sb
}

// ToStrings performs a quoted split to a slice of strings.
func ToStrings(b []byte) []string {
	var ss []string
	cp := Locations(b) // get chunk positions
	for _, pos := range cp {
		ss = append(ss, string(b[pos[0]:pos[1]]))
	}
	return ss
}

// ToStringBytes performs a quoted split, returning the first chunk as
// a string and the rest as a slice of byteslices.
func ToStringBytes(b []byte) (string, [][]byte) {
	bslices := ToBytes(b)
	return string(bslices[0]), bslices[1:]
}

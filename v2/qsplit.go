/*
Package qsplit (short for "quoted split") performs a Unix shell style
split-on-whitespace of its input. Its functions return the
non-whitespace "chunks" contained in their input, treating text within
balanced quotes as a single chunk.

Whitespace, according to qsplit, is the ASCII space and horizontal tab
characters.

By default, qsplit is aware only of the ASCII single and double quote
characters as chunk delineators. This can be changed with `SetQuotes`.

These are the rules used to delineate quoted chunks:

    - Quotes begin only at a word boundary
    - Quotes extend to the first closing quotation mark which matches the
      opening quote, which may or may not be at a word boundary.
    - Quotes do not nest

*/
package qsplit // import "github.com/firepear/qsplit/v2"

// Copyright (c) 2014-2020 Shawn Boyette <shawn@firepear.net>. All
// rights reserved.  Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"unicode/utf8"
)

var (
	// the quotation marks we know about
	qo = []rune{'\'', '"'}
	qc = []rune{'\'', '"'}
	// length of the list of quotes
	qlen = len(qo)
	// this gets returned when LocationsOnce finds no chunks
	onceNoLocs = [3]int{-1, -1, -1}
)


// SetQuotes sets the list of runes which will be considered
// quote-open and quote-close characters. As an example, to emulate
// the behavior of old versions of qsplit, the call would be:
//
//    SetQuotes([]rune{'\'', '"', '`', '‹', '«', '「', '『'},
//              []rune{'\'', '"', '`', '›', '»', '」', '』'})
func SetQuotes(qopen, qclose []rune) {
	qo = qopen
	qc = qclose
	qlen = len(qo)
}

// Locations returns the beginning and end points of all text chunks
// in its input.
func Locations(b []byte) [][2]int {
	return realLocations(b, false)
}

// LocationsOnce returns the beginning and end point of the first
// chunk, and the beginning of the next chunk. If this is all you
// need, LocationsOnce is significantly faster than Locations.
//
// If no chunks are found, the 1st element of the returned array will
// be -1. If only one chunk is found, the 3rd element will be -1.
func LocationsOnce(b []byte) [3]int {
	s := realLocations(b, true)
	slen := len(s)
	if slen == 2 {
		return [3]int{ s[0][0], s[0][1], s[1][0] }
	} else if slen == 1 {
		return [3]int{ s[0][0], s[0][1], -1 }
	} else {
		return onceNoLocs
	}
}

// realLocations does the work for Locations and LocationsOnce
func realLocations(b []byte, once bool) [][2]int {
	var si [][2]int     // slice of tuples of ints (chunk locations)
	var inw, inq bool   // in-word, in-quote flags
	var rune, endq rune // current rune; end-quote for current quote
	var i, idx int      // first index of chunk; byte index of current rune
	var j int

	// we need to operate at the runes level
	runes := bytes.Runes(b)
	for _, rune = range runes {
		switch {
		case inq:
			// in a quoted chunk, if we're looking at the
			// ending quote, unset inq and append a the
			// tuple for this chunk to the return list.
			if rune == endq {
				inq = false
				si = append(si, [2]int{i, idx})
			}
		case rune == ' ' || rune == '\t':
			// if looking whitespace and inw is set, end
			// the present chunk and append a new
			// tuple. else just move on.
			if inw {
				inw = false
				si = append(si, [2]int{i, idx})
			}
		case inw:
			// if in a regular word, do nothing
		default:
			// loop over quote-open runes, looking for a
			// match
			for j = 0; j < qlen; j++ {
				if rune == qo[j] {
					// looking at a quote; set
					// endq, inq, and i
					endq = qc[j]
					inq = true
					// if this is the 2nd chunk and we're
					// in once mode, return now
					i = idx + utf8.RuneLen(rune)
					if once && len(si) == 1 && ( inq || inw) {
						si = append(si, [2]int{i, -1})
						return si
					}
					// quit checking qos
					break
				}
			}
			if !inq {
				// not a quote, so we're looking at
				// the first rune in a word. set inw &
				// i
				inw = true
				i = idx
				// if this is the 2nd chunk and we're
				// in once mode, return now
				if once && len(si) == 1 && ( inq || inw) {
					si = append(si, [2]int{i, -1})
					return si
				}
			}
		}
		// else update idx and prune
		idx += utf8.RuneLen(rune)
	}
	// append the tuple for the last chunk if we were still in a
	// word or quote
	if inw || inq {
		si = append(si, [2]int{i, idx})
	}
	return si
}

// ToBytes performs a quoted split to a slice of byteslices.
func ToBytes(b []byte) [][]byte {
	var sb [][]byte    // slice of slice of bytes
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

// Once performs a single quoted split, returning the first chunk
// found in the input byteslice, and the remainder of the byteslice
func Once(b []byte) [][]byte {
	var sb [][]byte    // slice of slice of bytes
	cp := LocationsOnce(b) // get chunk positions
	if cp[2] == -1 {
		sb = append(sb, b)
	} else {
		sb = append(sb, b[cp[0]:cp[1]])
		sb = append(sb, b[cp[2]:])
	}
	return sb
}

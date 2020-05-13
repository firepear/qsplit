
# qsplit

Quoted splitting for Golang.


Qsplit splits a byte-slice into its constituent non-whitespace chunks,
keeping quoted chunks together, in the manner of the shell (mostly).

```
"foo bar baz"      -> [ b'foo', b'bar', b'baz'}
"   foo \tbar baz" -> [ b'foo', b'bar', b'baz'}
"'foo bar' baz"    -> [ b'foo bar', b`baz'}
"a b ‹c d "e f"›"  -> [ b'a', b'b', b'c d "e f"'}
"a b'cd e'f"       -> [ b'a', b"b'cd", b"e'f"}
```

See the [package doc](http://godoc.org/github.com/firepear/qsplit) for more
information.

## News

- 2020-05-13: v2.3.0: `qsplit` is now a Go module
- 2019-02-25: v2.2.3: Import path changed to github
- 2016-03-30: v2.2.2: `LocationsOnce` return values are more
  consistent when chunks are not found
- 2016-03-30: v2.2.1: `LocationsOnce` refactor resulting in a 2X speedup

## Use

`go get github.com/firepear/qsplit` or just import it!

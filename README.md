
# qsplit
Quoted string splitting for Golang.

NB: This package is in maintenance mode. There will be no development
in the forseeable future, except in response to issues/bug reports.

Qsplit turns a byte-slice into its constituent non-whitespace chunks
of bytes, keeping quoted chunks together, in the manner of the
shell. Mostly.

```
"foo bar baz"      -> [ b`foo`, b`bar`, b`baz` ]
"   foo \tbar baz" -> [ b`foo`, b`bar`, b`baz` ]
"'foo bar' baz"    -> [ b`foo bar`, b`baz` ]
"a b'cd e'f"       -> [ b`a`, b`b'cd`, b`e'f` ]
```
The library is tuned for speed, and the definition of quote characters
is configurable.

See the [package doc](https://pkg.go.dev/github.com/firepear/qsplit/v2?tab=doc) for more
information.

## News

- 2022-07-24: qsplit shifts to mainenance mode
- 2020-05-17: v2.5.0: Refactor of `LocationsOnce`; 1.1% speedup of
  `Locations`. `Once` now calls `LocationsOnce`
- 2020-05-16: v2.4.0: Refactor for speed improvements (min 1.5%;
  scales with input size). Default quote set is now single and double
  quote
- 2020-05-13: v2.3.1: Fix for import path after module transition
- 2020-05-13: v2.3.0: `qsplit` is now a Go module

See the [Release
notes](https://github.com/firepear/qsplit/blob/master/RELEASE_NOTES)
for all updates.

## Use

`go get github.com/firepear/qsplit/v2` or just import it!

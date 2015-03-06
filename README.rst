***************************
qsplit
***************************
Quoted splitting for Golang
===========================

Qsplit splits a byte-slice into its constituent non-whitespace chunks,
keeping quoted chunks together, in the manner of the shell.

::
   
    `foo bar baz`      -> {`foo`, `bar`, `baz`}
    `   foo \tbar baz` -> {`foo`, `bar`, `baz`}
    `'foo bar' baz`    -> {`foo bar`, `baz`}
    `a b ‹c d "e f"›`  -> {`a`, `b`, `c d "e f"`}
    `a b'cd e'f`       -> {`a`, `b'cd`, `e'f`}

See the package doc for more information.

* Current version: 2.1.0 (2015-02-15)

* Install: :code:`go get firepear.net/qsplit`

* `Release notes <http://github.com/firepear/qsplit/blob/master/RELEASE_NOTES>`_

* `Package documentation <http://godoc.org/firepear.net/qsplit>`_

* `Coverage report <http://firepear.net/qsplit/coverage.html>`_

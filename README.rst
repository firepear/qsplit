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

* Current version: 2.0.1 (2015-01-10)

* Install: :code:`go get firepear.net/goutils/qsplit`

* `Release notes <http://firepear.net/goutils/qsplit/RELEASE_NOTES.txt>`_

* `Package documentation <http://firepear.net:6060/pkg/firepear.net/goutils/qsplit/>`_

* `Coverage report <http://firepear.net/goutils/qsplit/coverage.html>`_

* `Issue tracker <https://firepear.atlassian.net/browse/QSPLIT>`_

* Source repo: :code:`git://firepear.net/goutils/qsplit.git`


Send questions, suggestions, or problem reports to shawn@firepear.net

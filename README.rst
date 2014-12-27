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
    `a b'cd e'f`       -> {`a`, `b'cd`, `e'f`} // open quote must be on word boundary

See the package doc for more information.
    
* Repository: :code:`git://firepear.net/goutils/qsplit.git`

* `Coverage report <http://firepear.net/goutils/qsplit/coverage.html>`_

Send questions, suggestions, or problem reports to shawn@firepear.net

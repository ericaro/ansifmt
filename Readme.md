[![Build Status](https://travis-ci.org/ericaro/ansifmt.png?branch=master)](https://travis-ci.org/ericaro/ansifmt) [![GoDoc](https://godoc.org/github.com/ericaro/ansifmt?status.svg)](https://godoc.org/github.com/ericaro/ansifmt)


# ansifmt library

`ansifmt` is a library to play around with [ansi escape codes](http://en.wikipedia.org/wiki/ANSI_escape_code).

It provides, mainly, a `Format` object describing most of the ansi escape codes capability (color, bold, etc.), and the ability to generate a "Coder" from it.

    var f Format
    f.SetStrike(true)
    striker := f.Coder()
    //striker is a function that will, decorate any word with strike code 
    fmt.Printf("%q", striker("Hello"))
    //Output: "\x1b[9mHello\x1b[29m"

It also provides a Writer, where you can change the output format.

# ansifmt/renderer 

it's a renderer for the excellent [blackfriday](http://github.com/russross/blackfriday) mardown generator.

Used in 'blackfriday', you can pretty print Markdown to the terminal console.

# License

ansifmt is available under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).

# Branches

master: [![Build Status](https://travis-ci.org/ericaro/ansifmt.png?branch=master)](https://travis-ci.org/ericaro/ansifmt) against go versions:

  - 1.2
  - 1.3
  - tip

dev: [![Build Status](https://travis-ci.org/ericaro/ansifmt.png?branch=dev)](https://travis-ci.org/ericaro/ansifmt) against go versions:

  - 1.2
  - 1.3
  - tip






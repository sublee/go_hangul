
[install go]: http://golang.org/install.html "Install Go"
[the godoc url]: http://localhost:6060/pkg/github.com/suapapa/hangul/ "the Godoc URL"

# About hangul

![logo](https://lh5.googleusercontent.com/-yblxhHfOiXw/UAzP9_3B0FI/AAAAAAAAA74/0nKCplLb9Ck/s615/IMG_20120723_131321-1.jpg)

hangul is a set of handy tools for manipulate korean character which can:

- Convert between jamo and compatibility-jamo.
- Split a character to it's three elements.
- Stroke count of a jamo.

## TODO

- Romanize.
- Lexycal Analize.


# Documentation

## Prerequisites

[Install Go][].

## Installation

    $ go get github.com/suapapa/hangul

## General Documentation

Use `go doc` to vew the documentation for hangul

    go doc github.com/suapapa/hangul

Or alternatively, use a godoc http server

    godoc -http=:6060

and visit [the Godoc URL][]


# Author

Homin Lee &lt;homin.lee@suapapa.net&gt;

# Copyright & License

Copyright (c) 2012, Homin Lee.
All rights reserved.
Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.

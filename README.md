go-walk
=======

<p align="center">
  <img src="https://github.com/while-loop/go-walk/blob/master/resources/gopherwalk.png">
  <br><br><br>
  <a href="https://godoc.org/github.com/while-loop/go-walk/walk"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
  <a href="https://travis-ci.org/while-loop/go-walk"><img src="https://img.shields.io/travis/while-loop/go-walk.svg?style=flat-square"></a>
  <a href="https://github.com/while-loop/go-walk/releases"><img src="https://img.shields.io/github/release/while-loop/go-walk.svg?style=flat-square"></a>
  <a href="https://coveralls.io/github/while-loop/go-walk"><img src="https://img.shields.io/coveralls/while-loop/go-walk.svg?style=flat-square"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-GPL%20v2-blue.svg?style=flat-square"></a>
</p>

Random Walk package written in Go.

Given weights, receive callbacks for
each direction.

Installation
------------

```
$ go get github.com/while-loop/go-walk/walk
```

Usage
-----

```go
package main

import (
	"github.com/while-loop/go-walk/walk"
	"fmt"
)

// Our walker that implements walk.Walker
type MyWalker struct {
	// keep a counter for each direction
	l, r, u, d int
}

func (w *MyWalker) Left() {
	w.l++
}

func (w *MyWalker) Right() {
	w.r++
}

func (w *MyWalker) Up() {
	w.u++
}

func (w *MyWalker) Down() {
	w.d++
}

func main() {
	mw := &MyWalker{}

	// Give custom weights to each direction
	// Left, Right, Up, Down
	w := walk.NewRandomWalk(10, 20, 30, 100, mw)

	// perform a walk with 32 iterations
	w.Walk(32)

	 // print the total amout of hits for each direction
	fmt.Println("My Walker:", *mw)
}

```

Changelog
---------

The format is based on [Keep a Changelog](http://keepachangelog.com/) 
and this project adheres to [Semantic Versioning](http://semver.org/).

[CHANGELOG.md](CHANGELOG.md)

License
-------
go-walk is licensed under the GPLv2 license. See [LICENSE](LICENSE) for details.

Author
------

Anthony Alves

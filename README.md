go-walk
=======

<p align="center">
  <img src="https://github.com/while-loop/go-walk/blob/master/resources/logo.png">
  <br><br><br>
  <a href="https://godoc.org/github.com/while-loop/go-walk/walk"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
  <a href="https://travis-ci.org/while-loop/go-walk"><img src="https://img.shields.io/travis/while-loop/go-walk.svg?style=flat-square"></a>
  <a href="https://github.com/while-loop/go-walk/releases"><img src="https://img.shields.io/github/release/while-loop/go-walk.svg?style=flat-square"></a>
  <a href="https://coveralls.io/github/while-loop/go-walk"><img src="https://img.shields.io/coveralls/while-loop/go-walk.svg?style=flat-square"></a>
  <a href="LICENSE"><img src="https://img.shields.io/github/license/mashape/apistatus.svg?style=flat-square"></a>
</p>

Random Walk package written in Go


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

type MyWalker struct {
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
	w := walk.NewRandomWalk(10, 20, 30, 100, mw)
	w.Walk(32)
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
go-walk is licensed under the MIT license. See [LICENSE](LICENSE) for details.

Author
------

Anthony Alves
# Section 26: Documentation

`godoc` extracts documentation from code and generates documentation for Packages.

`go doc` shows documentation for local and built in code. Use it for eg. viewing
the documentation for your package.

`go doc` is mostly used for programmatic handling of documentation, and `godoc`
for humans.

## Writing documentation in Go

The documentation of Go code happens in the code as comments. By tying code
and documentation together, we're always able to get explanations for how the
code is supposed to work and be used, right next to the actual code.

The rules (for generated documentation) are

* Write in complete sentences, begin with the name of the element being documented.
* Documentation is only generated for public symbols
* When documenting a package, you can use an `doc.go` file if you have a lot to
  say.

```go
// Package stuff - a package filled with stuffy stuff.
package stuff

import "fmt"

// DoIt does whatever needs doing.
// You really should try it, it's awesome!
func DoIt() {
	fmt.Println("Doing it")
}
```

Use a `doc.go` to document a package in detail.

```go
/*
Package stuff is actually a bit more

I have a lot to say about this package.

A LOT.

But It'll have to way for another day.
*/
package stuff
```

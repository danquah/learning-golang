# Section 6 - Programming fundamentals

A section that goes through the basics of how a computer represents data, and
then teaches about the different variable types.

I'm going to fast-forward through most of this. Which is going to be tricky
as there are some points in here that are good for me to re-iterate, but there
is also just a lot of things I already know.

The section on numeral systems is pretty good, Todd does a nice introduction to
decimal, binary and hexadecimals.

**Constants** can be "Constants of a kind" like so:

```go
const a = 42
```

And typed constants

```go
const a int64 = 42
```

Constants of a kind allows the compiler a bit more flexibility as to how the
value is stored.

Iota gives you a scoped monotonically incrementing counter. Should be used
carefully as it is a bit opaque for the average programmer.

```go

const (
  a = iota
  b
  c
)
// New const scope, so iota resets
const (
  d = iota
  e
  f
)

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)
fmt.Println(e)
fmt.Println(f)

// Prints
/*
0
1
2
0
1
2
*/

```

package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Not possible as the types are different
	// a = b

	// Conversion ("cast", but we do not use that word in Go)
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}

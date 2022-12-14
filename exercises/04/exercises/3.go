package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true

	// Print out the zero-values
	s := fmt.Sprintln(x, y, z)
	fmt.Println(s)

}

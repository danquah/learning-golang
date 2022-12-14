package main

import (
	"fmt"
)

type moo int

var y int
var x moo

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 43
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

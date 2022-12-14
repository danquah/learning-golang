package main

import (
	"fmt"
)

type moo int

func main() {
	var x moo

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}

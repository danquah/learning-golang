package main

import (
	"fmt"
)

var z int

func bar() {
	fmt.Println(z)
}

var y = 43

func foo() {
	fmt.Println(y)
}

func main() {
	// Short declaration operator
	x := 42
	fmt.Println(x)

	foo()

	bar()

}

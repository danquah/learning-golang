package main

import (
	"fmt"
)

func callThis(f func(int) int) {
	fmt.Println("Called it and got ", f(42))
}

func main() {
	/*
		A “callback” is when we pass a func into a func as an argument. For this exercise,
		● pass a func into a func as an argument
	*/

	myAnonymousFunction := func(value int) int {
		return value + 42
	}

	callThis(myAnonymousFunction)
}

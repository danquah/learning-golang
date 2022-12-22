package main

import (
	"fmt"
)

func main() {
	/*
		● Build and use an anonymous func
	*/

	func() {
		fmt.Println("I'm anonymous")
	}()

}

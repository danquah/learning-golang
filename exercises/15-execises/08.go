package main

import (
	"fmt"
)

func maFactory() func(int) int {
	return func(raise int) int {
		return raise * raise
	}
}

func main() {
	/*
		 Create a func which returns a func
		● assign the returned func to a variable
		● call the returned func

	*/

	whatIsThis := maFactory()

	fmt.Println(whatIsThis(2))
	fmt.Println(whatIsThis(100))
}

package main

import (
	"fmt"
)

func main() {
	/*
		● Assign a func to a variable, then call that func
	*/

	whatIsThis := func(counter int) {
		fmt.Println("I'm anonymous ", counter)
	}

	defer whatIsThis(1)
	whatIsThis(2)
}

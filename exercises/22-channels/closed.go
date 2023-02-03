package main

import (
	"fmt"
)

func main() {
	// Setup an unbuffered channel
	c := make(chan int)

	// Read data off the channel
	go func() {
		fmt.Println(<-c)
	}()

	// Send the value to be read
	c <- 42

	close(c)

	// Reads of a closed channel returns immediately so we can safely do it from
	// main
	v, ok := <-c

	// Outputs 0 false
	fmt.Println(v, ok)
}

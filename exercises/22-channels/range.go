package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int)
	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	// Signal that we're done and the reader can stop.
	close(c)
}

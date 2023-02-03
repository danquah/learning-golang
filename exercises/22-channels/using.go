package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	// Unbufferede, so both parties needs to be present to exchange data.
	wg.Add(2)
	go foo(c)
	go bar(c)
	wg.Wait()

	// Buffered, so we can do it async.
	c2 := make(chan int, 1)

	foo2(c2)
	bar2(c2)

}

func foo(c chan<- int) {
	c <- 42
	wg.Done()
}

func bar(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}

func foo2(c chan<- int) {
	c <- 42
}

func bar2(c <-chan int) {
	fmt.Println(<-c)
}

package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// Unidirectional
	chanUni := make(chan int)
	// This we can only write to
	chanRead := make(chan<- int)
	// This we can only Read from
	chanSend := make(<-chan int)

	// We can assign uni chanells to read/write
	chanRead = chanUni
	chanWrite = chanUni

}

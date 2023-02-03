package main

import (
	"fmt"
	"sync"
)

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar: ", i)
	}
	wg.Done()

}

var wg sync.WaitGroup

/*
in addition to the main goroutine, launch two additional goroutines
- each additional goroutine should print something out
- use waitgroups to make sure each goroutine finishes before your program exists
*/
func main() {
	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
	fmt.Println("Done")

}

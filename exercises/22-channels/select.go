package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("Add we're done")

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 42
	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v, ok := <-e:
			if !ok {
				continue
			}
			fmt.Println("Even: ", v)
		case v, ok := <-o:
			if !ok {
				continue
			}
			fmt.Println("Odd: ", v)
		case <-q:
			fmt.Println("We're done")
			return
		}
	}
}

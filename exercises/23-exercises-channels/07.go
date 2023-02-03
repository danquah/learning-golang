package main

import "fmt"

func main() {

	var gens []chan int

	for i := 0; i < 10; i++ {
		c := make(chan int, 10)
		go gen(c)
		gens = append(gens, c)

	}

	for _, c := range gens {
		for v := range c {
			fmt.Println(v)
		}
	}
}

func gen(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

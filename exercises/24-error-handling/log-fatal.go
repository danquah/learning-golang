package main

import "log"

func main() {
	log.Println("Hello")
	defer foo()

	log.Fatalln("Fatal")
}

func foo() {
	// This won't be called due to fatal.
	log.Println("I'm foo")
}

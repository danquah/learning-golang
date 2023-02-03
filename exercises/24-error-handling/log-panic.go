package main

import "log"

func main() {
	log.Println("Hello")
	defer foo()

	log.Panicln("Panic")
}

func foo() {
	// This will be called.
	log.Println("I'm foo")
}

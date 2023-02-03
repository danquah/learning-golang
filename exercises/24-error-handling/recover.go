package main

import "log"

func main() {
	log.Println("Hello")
	defer foo()

	log.Panicln("Panic")

	log.Println("I'll be called if we recover")
}

func foo() {
	log.Println("I'm foo")

	// Are we panicking?
	if i := recover(); i != nil {
		log.Println("Recovering from ", i)
	}
}

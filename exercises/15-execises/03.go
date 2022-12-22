package main

import "fmt"

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}

func main() {
	// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

	defer foo()
	bar()

}

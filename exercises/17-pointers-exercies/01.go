package main

import "fmt"

/*
● Create a value and assign it to a variable.
● Print the address of that value.
*/
func main() {
	myVar := "moo"
	fmt.Printf("%v\n", myVar)
	fmt.Printf("%v\n", &myVar)

}

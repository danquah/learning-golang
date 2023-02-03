package main

import "fmt"

/*
● create a person struct
● create a func called “changeMe” with a *person as a parameter
	○ change the value stored at the *person address
● “As an exception, if the type of x is a named pointer type and (*x).f
is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
○ https://golang.org/ref/spec#Selectors ● create a value of type person
○ print out the value ● in func main
○ call “changeMe” ● in func main
○ print out the value
*/

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "new name"
}

func main() {
	myPerson := person{"Ole"}
	fmt.Println(myPerson.name)
	changeMe(&myPerson)
	fmt.Println(myPerson.name)

}

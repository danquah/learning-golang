package main

import (
	"fmt"
)

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hello, my name is ", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

/*
This exercise will reinforce our understanding of method sets:
- create a type person struct
- attach a method speak to type person using a pointer receiver
  - *person

- create a type human interface
  - to implicitly implement the interface, a human must have the speak method

- create func “saySomething”
  - have it take in a human as a parameter
  - have it call the speak method

- show the following in your code
  - you CAN pass a value of type *person into saySomething
  - you CANNOT pass a value of type person into saySomething
*/
func main() {
	myHuman := person{"ole"}

	// breaks
	// saySomething(myHuman)
	saySomething(human(&myHuman))
	saySomething(&myHuman)
}

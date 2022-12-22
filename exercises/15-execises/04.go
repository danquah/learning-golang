package main

import "fmt"

func (p person) speak() {
	fmt.Printf("My name is %v %v and I am %d years old", p.first, p.last, p.age)
}

type person struct {
	first string
	last  string
	age   int
}

func main() {
	/*
		● Create a user defined struct with
		○ the identifier “person”
		○ the fields:
		■ first ■ last ■ age
		● attach a method to type person with
		○ the identifier “speak”
		○ the method should have the person say their name and age
		● create a value of type person
		● call the method from the value of type person
	*/

	myPerson := person{
		first: "Bjarne",
		last:  "Hansen",
		age:   42,
	}
	myPerson.speak()

}

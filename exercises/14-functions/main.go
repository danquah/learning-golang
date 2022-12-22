package main

import (
	"fmt"
)

func simpleFunction() {
	fmt.Println("that was easy")
}

func multiReturn() (bool, string) {
	return false, "muggie"
}

func variadicFunction(s ...string) {
	fmt.Printf("Slice has length: %v, capacity: %v, data: %v\n", len(s), cap(s), s)

	for i, v := range s {
		fmt.Println(i, v)
	}
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}

func useDefer() {
	defer foo()
	bar()
}

type PersonName struct {
	first string
	last  string
}

type Animal struct {
	name string
}

type Human struct {
	first string
	last  string
}

func (a Animal) speak(message string) {
	fmt.Printf("%v says: %v\n", a.name, message)
}

func (h Human) speak(message string) {
	fmt.Printf("%v %v says: %v\n", h.first, h.last, message)
}

type Speaker interface {
	speak(string)
}

func react(speaker Speaker, message string) {
	speaker.speak(message)
}

func (p PersonName) greet(say string) {
	fmt.Printf("%v %v %v\n", say, p.first, p.last)
}

// Just as an example we'll have the function both print and return the number.
func makeANumberCaller() func(int) int {
	return func(num int) int {
		fmt.Printf("The number is %d", num)
		return num
	}
}

func returnEnclosed() func() {
	inTheEnclosedScope := "Some value"

	return func() {
		fmt.Printf("The enclosed value is %v", inTheEnclosedScope)
	}
}

func main() {
	simpleFunction()
	multi1, multi2 := multiReturn()
	fmt.Println(multi1, multi2)
	variadicFunction("en", "lille", "and", "med", "vinger")
	variadicFunction()
	maSlice := []string{"boom", "boom", "boom"}
	variadicFunction(maSlice...)

	person1 := PersonName{first: "Some", last: "Person"}
	person1.greet("Hello")

	human1 := Human{first: "John", last: "Doe"}
	animal1 := Animal{name: "GoodBoy"}

	var speaker Speaker = human1
	react(human1, "I come in peace")
	react(animal1, "Woof!")

	// Type assertion
	fmt.Println(speaker.(Human).first)
	// The following would panic
	// fmt.Println(speaker.(Animal).first)

	// A safer way is to switch on the type:
	switch speaker := speaker.(type) {
	case Animal:
		fmt.Println("The speaker is an animal named ", speaker.name)
	case Human:
		fmt.Println("The speaker is a human named ", speaker.first)
	default:
		fmt.Println("No idea which typer of speaker this is")
	}

	// We can use a function as a value.
	myFunc := func() {
		fmt.Println("Im the first anonymous")
	}

	// Or we can run the function right away
	func() {
		fmt.Println("I'm the second anonymous")
	}()

	// We can run the stored function like so
	myFunc()
	myFunc()

	// Get a func from a func
	numberCaller := makeANumberCaller()
	numberCaller(42)

	{
		anEnclosedVar := "secret stuff"
		fmt.Println(anEnclosedVar)
	}

	// This would not compile.
	// fmt.Println(anEnclosedVar)

	// Run a function that has enclosed a value
	returnEnclosed()()
}

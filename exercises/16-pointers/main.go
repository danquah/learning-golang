package main

import (
	"fmt"
)

func main() {
	// Declare and initialize a variable
	a := 42
	fmt.Println(a)
	// Get the address for the variable
	fmt.Println(&a)

	// Verify the types of the things we just printed
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// Now, store the pointer(address) to the a variable in a new variable.
	b := &a
	// Still a pointer.
	fmt.Println(b)
	// Now dereference the value of b, and we should end up with the value of a
	fmt.Println(*b)

	// We can also do a very convoluted assignment to a by doing
	*b = 44
	fmt.Println(a)

	// More compact, a silly way of copying a variable.
	myVar := "some string"
	myPointer := &myVar
	fmt.Println(myPointer)
	copyOfMyVar := *myPointer
	fmt.Println(copyOfMyVar)

	// pass a variable by reference to a function and have it mutate the callers
	// instance.
	moo := 41
	foo(&moo)
	fmt.Println(moo)

	// the selector auto dereferences struct pointers
	someStruct := myStruct{"someSting"}
	// "normal" access to the field
	fmt.Println(someStruct.dummy)
	someStructReference := &someStruct
	// Access via the pointer - same syntax.
	fmt.Println(someStructReference.dummy)

	// The same goes for arrays
	myArr := [2]string{"foo", "bar"}
	myArrPointer := &myArr
	fmt.Println(myArr[1])
	fmt.Println(myArrPointer[1])

	// A receiver can require a pointer to the value your are passing.
	// In the following case the interface itself does not specify anything about
	// pointers, but one of the attached method references it's receiver by pointer
	// and thus the value passed in must be "by reference".
	callByPointer(&someStruct)
	callByValue(someStruct)

}

func foo(y *int) {
	*y = *y + 1
}

type myStruct struct {
	dummy string
}

func callByPointer(thing hasPointerReceiver) {
	thing.takesPointer()
}

func callByValue(thing hasValueReceiver) {
	thing.takesValue()
}

func (s *myStruct) takesPointer() {
	s.dummy = "foo"
	fmt.Printf("%T\n", s)
	fmt.Println("I take a pointer")
}

func (s myStruct) takesValue() {
	fmt.Printf("%T\n", s)
	fmt.Println("I take a value")
}

type hasPointerReceiver interface {
	takesPointer()
	takesValue()
}

type hasValueReceiver interface {
	takesValue()
}

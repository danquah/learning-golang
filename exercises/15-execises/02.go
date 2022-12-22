package main

import "fmt"

func foo(v ...int) int {
	var sum int

	for _, v := range v {
		sum += v
	}

	return sum
}

func bar(v []int) int {
	var sum int

	for _, v := range v {
		sum += v
	}

	return sum
}

func main() {
	/*
		create a func with the identifier foo that
		○ takes in a variadic parameter of type int
		○ pass in a value of type []int into your func (unfurl the []int)
		○ returns the sum of all values of type int passed in
		● create a func with the identifier bar that
		○ takes in a parameter of type []int
		○ returns the sum of all values of type int passed in
	*/

	fmt.Println(foo([]int{1, 2, 3}...))
	fmt.Println(bar([]int{1, 2, 3}))

}

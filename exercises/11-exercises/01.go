package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	arr[0] = 5
	arr[1] = 10
	arr[2] = 15
	arr[3] = 20
	arr[4] = 25

	for x := range arr {
		fmt.Printf("%T, %v\n", x, x)
	}

	arr2 := [5]int{1, 2, 3, 4, 5}

	for x := range arr2 {
		fmt.Printf("%T, %v\n", x, x)
	}

}

package main

import (
	"fmt"
)

func main() {
	/*
	   ● Using a COMPOSITE LITERAL:
	   ● create a SLICE of TYPE int
	   ● assign 10 VALUES
	   ● Range over the slice and print the values out.
	   ● Using format printing
	   ○ print out the TYPE of the slice
	*/

	maSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range maSlice {
		fmt.Printf("%T %v\n", v, v)
	}
}

package main

import (
	"fmt"
)

func main() {
	start := 42
	end := 51
	offsets := []int{0, 5, 2, 1}
	length := 5

	source := []int{}

	for i := 0; i <= (end - start); i++ {
		source = append(source, start+i)
	}

	for _, offset := range offsets {
		fmt.Println(source[offset : offset+length])
	}

}

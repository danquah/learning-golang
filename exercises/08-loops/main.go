package main

import (
	"fmt"
)

func main() {
	fmt.Println("Them loops gona loop")

	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}

	j := 0
	for {
		j++
		if j > 10 {
			break
		}
		fmt.Println(j)

	}

	for k := 0; k < 12; k++ {
		if k%2 == 0 {
			continue
		}

		fmt.Println(k)
	}

	// We can have initializations in if-statements as well
	if x := 42; x > 10 {
		fmt.Println("yay")
	}

	x := 42
	switch {
	case x > 42:
		fmt.Println("Larger than")
	case x < 42:
		fmt.Println("Less than")
	case x == 42:
		fmt.Println("Equal to")
		// Never do this
		fallthrough
	case x == 41, x == 42, x == 43:
		fmt.Println("41 or 42 or 43")
	}

	y := 42

	switch y {
	case 41:
		fmt.Println(41)
	case 42:
		fmt.Println(42)
	}

}

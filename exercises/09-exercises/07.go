package main

import (
	"fmt"
)

func main() {
	x := 44
	if x < 10 {
		fmt.Println("Less")
	} else if x > 100 {
		fmt.Println("Quite large")
	} else {
		fmt.Println("Beats me")
	}
}

package main

import (
	"fmt"
)

func main() {
	multi := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, outer := range multi {
		for _, inner := range outer {
			fmt.Printf("%v", inner)
		}
	}
}

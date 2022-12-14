package main

import (
	"fmt"
)

func main() {
	/*
		Using the code from the previous example, delete a record from your map.
		Now print the map out using the “range” loop


	*/
	people := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	people["donald"] = []string{"sleep", "food"}

	delete(people, "bond_james")

	for k, favorites := range people {
		fmt.Printf("%v:\n", k)
		for i, v := range favorites {
			fmt.Printf("\t%v: %v\n", i, v)
		}
	}
}

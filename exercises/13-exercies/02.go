package main

import "fmt"

/*
Take the code from the previous exercise, then store the values of type person
in a map with the key of last name. Access each value in the map. Print out the
values, ranging over the slice.
*/
func main() {
	type person struct {
		first  string
		last   string
		favIce []string
	}

	person1 := person{
		first:  "Some",
		last:   "Dude",
		favIce: []string{"Vanilla"},
	}

	person2 := person{
		first:  "Sam",
		last:   "Samson",
		favIce: []string{"Strawberry"},
	}

	people := map[string]person{
		"Dude":   person1,
		"Samson": person2,
	}

	for k, v := range people {
		fmt.Println(k, v)
	}

}

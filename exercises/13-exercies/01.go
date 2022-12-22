package main

import "fmt"

/*
Create your own type “person” which will have an underlying type of “struct”
so that it can store the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
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

	for _, v := range person1.favIce {
		fmt.Println(person1, v)
	}

	for _, v := range person2.favIce {
		fmt.Println(person2, v)
	}

}

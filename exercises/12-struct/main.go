package main

import (
	"fmt"
)

func main() {

	type person struct {
		first string
		last  string
	}

	dude := person{
		first: "the",
		last:  "dude",
	}

	fmt.Println(dude)

	// Access / dereference using dot notation
	fmt.Printf("He is called %v %v", dude.first, dude.last)

	type importantPerson struct {
		person
		role string
	}

	santa := importantPerson{
		person: person{
			first: "Santa",
			last:  "Claus",
		},
		role: "Distribute presents",
	}

	fmt.Println(santa)

	fmt.Println(santa.first)
	fmt.Println(santa.person.first)

	type twins struct {
		person1 person
		person2 person
		role    string
	}

	// Anonymous / unnamed fields gets their names after their types
	type silly struct {
		string
	}
	somePerson := struct {
		first string
		last  string
	}{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(somePerson)

}

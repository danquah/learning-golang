package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Ole",
		Last:  "Hansen",
		Age:   42,
	}
	p2 := person{"Jens", "Peter", 24}
	people := []person{p1, p2}
	fmt.Println(p1)
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Print("Error: ", err)
	}
	// We now have a json string.
	fmt.Println(string(bs))

	// Unmarshal into a variable of type []person
	var peopleUnmarshal []person
	err = json.Unmarshal(bs, &peopleUnmarshal)
	if err != nil {
		fmt.Print("Error while unmarshaling: ", err)
	}

	fmt.Printf("%v\n", peopleUnmarshal)

	// Lets say we want the json to be lowercased we can then tag our struct
	type LowerCasedPerson struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Age   int    `json:"age"`
	}
	jsonString := `[{"first":"Jens","last":"Peter","Age":24}]`
	var lowPeople []LowerCasedPerson
	err = json.Unmarshal([]byte(jsonString), &lowPeople)
	if err != nil {
		fmt.Print("Error while unmarshaling: ", err)
	}

	for i, v := range lowPeople {
		fmt.Printf("%v: %v\n", i, v)
	}
}

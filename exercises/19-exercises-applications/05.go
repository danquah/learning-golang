package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByUser []user

// Len implements sort.Interface
func (u ByUser) Len() int {
	return len(u)
}

// Less implements sort.Interface
func (u ByUser) Less(i int, j int) bool {
	return u[i].Age < u[j].Age || u[i].Last < u[j].Last
}

// Swap implements sort.Interface
func (u ByUser) Swap(i int, j int) {
	u[j], u[i] = u[i], u[j]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	sort.Sort(ByUser(users))

	for _, v := range users {
		sayings := v.Sayings
		sort.Strings(sayings)
		fmt.Printf("%v %v is %d years old and says\n", v.First, v.Last, v.Age)
		for _, saying := range sayings {
			fmt.Printf("- %v\n", saying)
		}
	}

}

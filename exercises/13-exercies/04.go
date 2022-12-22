package main

import "fmt"

/*
Take the code from the previous exercise, then store the values of type person
in a map with the key of last name. Access each value in the map. Print out the
values, ranging over the slice.
*/
func main() {
	somePerson := struct {
		first string
		last  string
	}{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(somePerson)

}

package main

import "fmt"

/*
Take the code from the previous exercise, then store the values of type person
in a map with the key of last name. Access each value in the map. Print out the
values, ranging over the slice.
*/
func main() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	myTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}

	mySedan := sedan{
		vehicle: vehicle{
			doors: 9,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(myTruck)
	fmt.Println(mySedan)
	fmt.Println(myTruck.fourWheel)
	fmt.Println(mySedan.luxury)
}

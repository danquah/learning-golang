package main

import (
	"fmt"
)

func main() {

	fmt.Println("Raw arrays")
	var x [5]int
	fmt.Println(len(x))
	x[2] = 99
	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println(len(x))

	y := []int{42, 41, 40}
	fmt.Println(y)

	fmt.Println("Slices")
	sli := []int{1, 2, 3, 8}

	// Check the length just like arrays.
	fmt.Println(len(sli))
	// The same with access via index
	fmt.Println(sli[2])

	// Loop over range
	// Notice range is a keyword, not a function.
	for index, value := range sli {
		fmt.Println(index, value)
	}

	slicySlice := []int{1, 2, 4, 8, 16}
	// Entire slices
	fmt.Println(slicySlice[0:])
	// Last element 2 elements, skipping the 3 first
	fmt.Println(slicySlice[3:])
	// Last element
	fmt.Println(slicySlice[4:5])

	// Silly iteration
	for i := 0; i < len(slicySlice); i++ {
		fmt.Println(i, slicySlice[i:i+1])
	}

	growingSlice := []int{1, 2, 3}
	growingSlice = append(growingSlice, 4, 5, 6)
	for _, v := range growingSlice {
		fmt.Println(v)
	}

	appendSlice := []int{7, 8, 9}
	// We can expand a slice into its elements and pass them to a varadic function
	growingSlice = append(growingSlice, appendSlice...)
	for _, v := range growingSlice {
		fmt.Println(v)
	}

	// Delete an element
	slice1 := []int{1, 2, 3}
	slice1 = append(slice1[0:1], slice1[2:]...)
	fmt.Println(slice1)

	mSlice1 := []string{"muggie", "moo"}
	mSlice2 := []string{"foo", "bar"}

	multiSlice := [][]string{mSlice1, mSlice2}
	fmt.Println(multiSlice)

	// Maps
	m := map[string]int{
		"tree":  10,
		"shrub": 4,
	}

	// Prints 0
	fmt.Println(m["mountain"])

	if v, ok := m["mountain"]; ok {
		fmt.Println("Only prints if the mountain was there ", v)
	}

	m["mountain"] = 1000
	// Use the init section of if statements to first check if the element is
	// there, then check the value of ok.
	if v, ok := m["mountain"]; ok {
		fmt.Println("Now it is actually there ", v)
	}

	// And then get rid of it again
	delete(m, "mountain")
	if v, ok := m["mountain"]; ok {
		fmt.Println("Won't print", v)
	}

	// Ranging over a map
	for k, v := range m {
		fmt.Println(k, v)
	}

}

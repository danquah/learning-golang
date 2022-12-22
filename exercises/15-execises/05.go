package main

import (
	"fmt"
	"math"
)

type square struct {
	length int
	width  int
}

type circle struct {
	radius int
}

type shape interface {
	area() int
}

func (c circle) area() int {
	return int(math.Pi * math.Pow(float64(c.radius), 2))
}

func (s square) area() int {
	return s.length * s.width
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {
	/*
		● create a type SQUARE
		● create a type CIRCLE
		● attach a method to each that calculates AREA and returns it
		○ circle area= π r^2
		○ squarearea=L*W
		● create a type SHAPE that defines an interface as anything that has the AREA method
		● create a func INFO which takes type shape and then prints the area
		● create a value of type square
		● create a value of type circle
		● use func info to print the area of square
		● use func info to print the area of circle
	*/

	mySquare := square{length: 5, width: 10}
	myCircle := circle{radius: 12}

	info(mySquare)
	info(myCircle)
}

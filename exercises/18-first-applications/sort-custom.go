package main

import (
	"fmt"
	"sort"
)

type Person struct {
	first string
	age   int
}

type ByAge []Person
type ByName []Person

func (p Person) String() string {
	return fmt.Sprintf("%v is %d years old\n", p.first, p.age)
}

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a ByAge) Swap(i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func (bn ByName) Len() int {
	return len(bn)
}

func (bn ByName) Less(i, j int) bool {
	return bn[i].first < bn[j].first
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}

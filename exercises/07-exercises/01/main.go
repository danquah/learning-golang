package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	n := 42
	fmt.Printf("%d %#x %b\n", n, n, n)

	e := 100

	f := 101

	eq := n == n
	lteq := n <= e
	gteq := n >= e
	neq := n != n
	lt := n < f
	gt := n > f

	fmt.Println(eq, lteq, gteq, neq, lt, gt)

	const c1 = 42
	const c2 int = 44

	fmt.Println(c1, c2)

	shift1 := 42
	fmt.Printf("%d %#x %b\n", shift1, shift1, shift1)

	shift2 := shift1 << 1

	fmt.Printf("%d %#x %b\n", shift2, shift2, shift2)

	str := `raaa
	w`

	fmt.Println(str)

	curYear, _ := strconv.ParseInt(time.Now().Format("2006"), 0, 64)

	fmt.Println("Current year", curYear)

	const (
		year1 = 2022 + iota
		year2
		year3
		year4
	)
	fmt.Println("Year ", year1)
	fmt.Println("Year ", year2)
	fmt.Println("Year ", year3)
	fmt.Println("Year ", year4)

}

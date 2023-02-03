package main

import (
	"fmt"
)

type thingy struct {
	value string
}

func main() {
	_, err := haveProblems()
	if err != nil {
		fmt.Print("Bad things happened: ", err)
	}
}

func haveProblems() (int, error) {
	return 0, thingy{"stuff happened"}
}

func (t thingy) Error() string {
	return fmt.Sprint("Custom error string for ", t.value)
}

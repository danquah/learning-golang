package main

import "fmt"

type customErr struct {
	message string
}

func (c customErr) Error() string {
	return c.message
}

func foo(err error) {
	fmt.Print("Got the error ", err)
}

func main() {
	foo(customErr{})
}

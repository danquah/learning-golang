package main

import (
	"fmt"
	"runtime"
)

/*
Create a program that prints out your OS and ARCH. Use the following commands to run it
- gorun
- go build
- go install
*/
func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

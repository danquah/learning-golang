package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("The boring way")
	fmt.Fprint(os.Stdout, "Much more 1337\n")
}

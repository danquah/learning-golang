package main

import (
	"fmt"
)

func main() {
	const start = int('A')
	for i := 0; i < 26; i++ {
		runeCode := i + start
		fmt.Println(runeCode)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t %#U \n", start+i)
		}
	}
}

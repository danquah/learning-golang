// Package dog has a lot of nifty dog-related functions.
package dog

import "fmt"

// Years can be used to convert from human to dog years.
func Years(humanYears int) int {
	return humanYears * 7
}

func main() {
	fmt.Println(Years(2))
}

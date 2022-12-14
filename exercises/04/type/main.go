package main

import (
	"fmt"
)

var y = 42
var z = "Muggie, Wookie"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	raw := `my raw
string`
	fmt.Println(raw)
	interpreted := "my interpreted string"
	fmt.Println(interpreted)
}

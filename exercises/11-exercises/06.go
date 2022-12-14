package main

import (
	"fmt"
)

func main() {
	source := []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}
	fmt.Printf("Source Length: %d, Capacity: %d\n", len(source), cap(source))
	dest := make([]string, len(source), len(source))

	fmt.Printf("Dest Length: %d, Capacity: %d\n", len(dest), cap(dest))
	for i, v := range source {
		dest[i] = v
	}
	fmt.Printf("After: dest Length: %d, Capacity: %d\n", len(dest), cap(dest))
}

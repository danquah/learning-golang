package main

import "testing"


func TestMyDoubler(t *testing.T) {
	input := 2
	expected := input * 2
	result := myDoubler(input)

	if result != expected {
		t.Error("Unexpected return ", result, " expected ", expected)
	}

}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Using goroutines, create an incrementer program
  - have a variable to hold the incrementer value
  - launch a bunch of goroutines
  - each goroutine should
  - read the incrementer value
  - store it in a new variable
  - yield the processor with runtime.Gosched()
  - increment the new variable
  - write the value in the new variable back to the incrementer
    variable

- use waitgroups to wait for all of your goroutines to finish
- the above will create a race condition.
- Prove that it is a race condition by using the -race flag
*/
var wg sync.WaitGroup

func main() {
	const numToLaunch = 50
	wg.Add(numToLaunch)
	incShared := 0
	for i := 0; i < numToLaunch; i++ {
		go func() {
			tmpInc := incShared
			fmt.Println("Before ", tmpInc)
			runtime.Gosched()
			incShared = tmpInc + 1
			fmt.Println("After ", incShared)
			wg.Done()
		}()
		fmt.Println("Launched ", i)
	}
	wg.Wait()
}

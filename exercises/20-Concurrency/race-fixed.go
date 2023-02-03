package main

import (
	"fmt"
	"runtime"
	"sync"
)

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar: ", i)
	}
}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("Arch\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Rutines\t\t", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// Yield
			runtime.Gosched()
			// Increment the shared variabel. This is not synchronized so we'll run
			// in to trouble as it may have increased while we where not scheduled.
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Gorutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Rutines\t\t", runtime.NumGoroutine())
	fmt.Println(counter)

}

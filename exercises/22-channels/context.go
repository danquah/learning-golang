package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Current err: ", ctx.Err())
	fmt.Println("Current gorutines", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return

			default:
				n++
				time.Sleep(time.Millisecond * 220)
				fmt.Println("working", n)
			}
		}
	}()

	duration := rand.Intn(30)
	fmt.Println("Current gorutines", runtime.NumGoroutine())
	fmt.Println("Sleeping for ", duration)
	time.Sleep(time.Second * time.Duration(duration))
	fmt.Println("Current gorutines", runtime.NumGoroutine())
	fmt.Println("Shutting down")
	cancel()
	time.Sleep(time.Second * time.Duration(5))
	fmt.Println("Current gorutines", runtime.NumGoroutine())
	fmt.Println("Exiting")
}

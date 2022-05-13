package main

import (
	"fmt"
	"sync"
	"time"
)

func fun(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	//fun("direct call")

	// TODO: write goroutine with different variants for function call.
	var wg sync.WaitGroup
	wg.Add(1)
	go fun("goroutine #1", &wg)

	wg.Add(1)
	go fun("goroutine #2", &wg)

	wg.Add(1)
	go fun("goroutine #3", &wg)

	// goroutine function call

	// goroutine with anonymous function
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fun("hello from anonymous", wg)
	}(&wg)

	// goroutine with function value call

	// wait for goroutines to end
	wg.Wait()
	fmt.Println("done..")
}

package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		c.L.Unlock()
	}()

	c.L.Lock()
	// writes changes to sharedRsc
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	c.Broadcast()
	c.L.Unlock()

	wg.Wait()
}

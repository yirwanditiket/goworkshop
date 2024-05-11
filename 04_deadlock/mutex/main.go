package main

import (
	"sync"
	"fmt"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	myVar := 0

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		// forgot to unlock
		myVar = 200
	}()

	wg.Wait()

	mu.Lock()
	fmt.Println(myVar)
	mu.Unlock()
}


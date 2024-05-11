package main

import (
	"fmt"
	"sync"
)

func main() {
	myFunc(0)
}

func myFunc(i int) {
	var mu sync.Mutex
	var wg sync.WaitGroup

	myVar := 0

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				fmt.Println("panic", r)
			}
		}()

		mu.Lock()
		myVar = 200
		myVar = myVar / i
		mu.Unlock()
	}()

	wg.Wait()

	fmt.Println("Waiting for mutex....")
	mu.Lock()
	fmt.Println(myVar)
	mu.Unlock()
}

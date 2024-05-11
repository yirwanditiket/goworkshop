package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 9; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("Hi")
		}()
	}

	wg.Wait()
}

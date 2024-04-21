package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	total := 0

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			fmt.Println("add -- #", i)
			total = total + 1
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println("Total: ", total)
}

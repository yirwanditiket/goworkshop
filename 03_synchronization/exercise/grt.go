package main

import (
	"fmt"
	"sync"
)

func grt(in []int) int {
	var mu sync.Mutex
	var wg sync.WaitGroup

	total := 0

	wg.Add(len(in))
	for _, i := range in {
		go func(i int) {
			defer wg.Done()

			x := fib(i)
			mu.Lock()
			total += x
			mu.Unlock()
			fmt.Println("Done", i)
		}(i)
	}

	wg.Wait()
	return total
}

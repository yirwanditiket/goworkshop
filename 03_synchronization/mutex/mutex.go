package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	total := 0

	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			fmt.Println("add -- #", i)
			total = total + 1
			mu.Unlock()
		}(i)
	}

	mu.Lock()
	fmt.Println("Total: ", total)
	mu.Unlock()
}

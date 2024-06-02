package main

import (
	"fmt"
	"sync"
	"time"
)

const BIG_NUMBER = 43

func main() {
	st := time.Now()

	var wg sync.WaitGroup
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)
		wg.Add(50)
		for i := 0; i < 50; i++ {
			go func() {
				defer func() {
					wg.Done()
				}()
				resultCh <- fib(BIG_NUMBER)
			}()
		}
		wg.Wait()
	}()

	total := 0
	for res := range resultCh {
		total += res
	}

	fmt.Println("Total", total)
	fmt.Println("Done, elapsed ", time.Since(st))
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

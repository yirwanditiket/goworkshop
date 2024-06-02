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

	// create semaphore with count "5"
	// using buffered channel
	sem := make(chan struct{}, 5)

	go func() {
		defer close(resultCh)

		wg.Add(50)
		for i := 0; i < 50; i++ {
			// acquire the semaphore by sending a value
			// if there is space in the buffer, then this operation shouldn't block
			// if the buffer is already full, then this operation will block
			sem <- struct{}{}

			go func() {
				defer func() {
					wg.Done()

					// release the semaphore by poping a value
					// this will make space in the buffer allowing another goroutine to acquire
					<-sem
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

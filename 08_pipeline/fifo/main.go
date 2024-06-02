package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10

	// sourceCh is receive only channel of integer
	sourceCh := sourcePipeline(n)

	// sourceCh become the upstream channel of the
	// square fan in fan out pipeline
	resultCh := squareFanInFanOut(sourceCh, 4)

	for res := range resultCh {
		fmt.Println(res)
	}
}

// sourcePipeline accept integer return channel of
// integer (receive only)
func sourcePipeline(n int) <-chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)

		for i := 0; i < n; i++ {
			outCh <- i
		}
	}()
	return outCh
}

// squarePipeline accept receive only channel of integer
// and return receive only channel of integer
func squarePipeline(inCh <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)

		// receive from the inbound channel
		for in := range inCh {
			// do some work
			res := in * in

			// send the result to outbound channel
			outCh <- res
		}
	}()
	return outCh
}

func squareFanInFanOut(inCh <-chan int, workerNum int) <-chan int {
	outCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(workerNum)

	sink := func(workerOutCh <-chan int) {
		defer wg.Done()
		// fan in
		for res := range workerOutCh {
			outCh <- res
		}
	}

	// fan out
	for i := 0; i < workerNum; i++ {
		go sink(squarePipeline(inCh))
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}

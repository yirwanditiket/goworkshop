package main

import (
	"fmt"
)

func main() {
	n := 10

	// sourceCh is receive only channel of integer
	sourceCh := sourcePipeline(n)

	// sourceCh become the upstream channel of the
	// square pipeline
	resultCh := squarePipeline(sourceCh)

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

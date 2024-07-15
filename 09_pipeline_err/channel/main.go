package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10

	// sourceCh is receive only channel of integer
	sourceCh, sourceErrCh := sourcePipeline(n)

	// sourceCh become the upstream channel of the
	// square fan in fan out pipeline
	resultCh, resultErrCh := squarePipeline(sourceCh)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for err := range mergeErrCh(sourceErrCh, resultErrCh) {
			fmt.Println("Encountered error ", err)
		}
	}()

	for res := range resultCh {
		fmt.Println(res)
	}

	wg.Wait()
}

// sourcePipeline accept integer return channel of
// integer (receive only)
func sourcePipeline(n int) (<-chan int, <-chan error) {
	outCh := make(chan int)
	errCh := make(chan error)
	go func() {
		defer func() {
			close(outCh)
			close(errCh)
		}()

		for i := 0; i < n; i++ {
			if i == 4 {
				errCh <- fmt.Errorf("Error while generating i = 4")
				continue
			}

			outCh <- i
		}
	}()
	return outCh, errCh
}

// squarePipeline accept receive only channel of integer
// and return receive only channel of integer
func squarePipeline(inCh <-chan int) (<-chan int, <-chan error) {
	outCh := make(chan int)
	errCh := make(chan error)
	go func() {
		defer func() {
			close(outCh)
			close(errCh)
		}()

		// receive from the inbound channel
		for in := range inCh {
			if in == 7 {
				errCh <- fmt.Errorf("error while squaring i = 7")
				continue
			}

			// do some work
			res := in * in

			// send the result to outbound channel
			outCh <- res
		}
	}()
	return outCh, errCh
}

func mergeErrCh(errChs ...<-chan error) <-chan error {
	outCh := make(chan error, len(errChs))

	var wg sync.WaitGroup
	wg.Add(len(errChs))

	sink := func(errCh <-chan error) {
		defer wg.Done()

		for err := range errCh {
			outCh <- err
		}
	}

	for _, errCh := range errChs {
		go sink(errCh)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}

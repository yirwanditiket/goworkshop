package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Encountered error ", err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	n := 10

	// sourceCh is receive only channel of integer
	sourceCh, sourceErrCh := sourcePipeline(ctx, n)

	// sourceCh become the upstream channel of the
	// square fan in fan out pipeline
	resultCh, resultErrCh := squarePipeline(ctx, sourceCh)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for res := range resultCh {
			fmt.Println(res)
		}
	}()

	for err := range mergeErrCh(sourceErrCh, resultErrCh) {
		if err != nil {
			return err
		}
	}

	wg.Wait()

	return nil
}

// sourcePipeline accept integer return channel of
// integer (receive only)
func sourcePipeline(ctx context.Context, n int) (<-chan int, <-chan error) {
	outCh := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		defer func() {
			close(outCh)
			close(errCh)
		}()

		for i := 0; i < n; i++ {
			if i == 4 {
				errCh <- fmt.Errorf("(Fatal) Error while generating i = 4")
				return
			}

			select {
			case <-ctx.Done():
				return
			case outCh <- i:
			}
		}
	}()
	return outCh, errCh
}

// squarePipeline accept receive only channel of integer
// and return receive only channel of integer
func squarePipeline(ctx context.Context, inCh <-chan int) (<-chan int, <-chan error) {
	outCh := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		defer func() {
			close(outCh)
			close(errCh)
		}()

		// receive from the inbound channel
		for in := range inCh {
			if in == 7 {
				errCh <- fmt.Errorf("((Fatal) error while squaring i = 7")
				return
			}

			// do some work
			res := in * in

			// send the result to outbound channel
			select {
			case <-ctx.Done():
				return
			case outCh <- res:
			}

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

package main

import (
	"context"
	"fmt"
)

type MyData struct {
	SourceI int
	SquareRes int
	Err error
}

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
	sourceCh := sourcePipeline(ctx, n)

	// sourceCh become the upstream channel of the
	// square pipeline
	resultCh := squarePipeline(ctx, sourceCh)

	for data := range resultCh {
		if data.Err != nil {
			fmt.Println("Encountered error ", data.Err)
		} else {
			fmt.Println(data.SquareRes)
		}
	}

	return nil
}

// sourcePipeline accept integer return channel of
// integer (receive only)
func sourcePipeline(ctx context.Context, n int) (<-chan MyData) {
	outCh := make(chan MyData)
	go func() {
		defer close(outCh)

		for i := 0; i < n; i++ {
			var err error
			sourceI := i
			if sourceI == 4 {
				sourceI = 0
				err = fmt.Errorf("(Fatal) Error while generating i = 4")
			}

			select {
			case <-ctx.Done():
				return
			case outCh <- MyData{
				SourceI: i,
				Err: err,
			}:
			}
		}
	}()
	return outCh
}

// squarePipeline accept receive only channel of integer
// and return receive only channel of integer
func squarePipeline(ctx context.Context, inCh <-chan MyData) (<-chan MyData) {
	outCh := make(chan MyData)
	go func() {
		defer close(outCh)

		// receive from the inbound channel
		for in := range inCh {
			if in.SourceI == 7 {
				in.Err = fmt.Errorf("((Fatal) error while squaring i = 7")
			}

			// do some work
			in.SquareRes = in.SourceI * in.SourceI

			// send the result to outbound channel
			select {
			case <-ctx.Done():
				return
			case outCh <- in:
			}

		}
	}()
	return outCh
}

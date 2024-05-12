package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	err := myFunc()
	fmt.Println("my func error", err)
	time.Sleep(2 * time.Second)
}

func myFunc() error {
	// creating channel to send and receive "int"
	// unbuffered, bi-directional
	ch := make(chan int)

	// cancel signal using context
	ctx, cancel := context.WithCancel(context.Background())

	// if this function exit, cancel the context
	// so any goroutine listening to this context can exit too
	defer cancel()

	go func() {
		// close the channel after there is no
		// more value to send
		defer func() {
			fmt.Println("The producer goroutine exit!")
			close(ch)
		}()

		for i := 0; i <= 2; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				fmt.Println("Exitting because of cancel")
				return
			}
		}
	}()

	time.Sleep(1 * time.Second)

	// this loop keep iterating for every value sent to
	// the channel, and done after the channel is closed
	for myVar := range ch {
		fmt.Println("Got value", myVar)
		if myVar == 1 {
			return fmt.Errorf("Can't handle 1, must exit")
		}
	}

	fmt.Println("All done")
	return nil
}

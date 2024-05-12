package main

import (
	"fmt"
)

func main() {
	// creating channel to send and receive "int"
	// unbuffered, bi-directional
	ch := make(chan int)

	go func() {
		// close the channel after there is no
		// more value to send
		defer close(ch)
		for i := 0; i <= 2; i++ {
			ch <- i
		}
	}()

	// this loop keep iterating for every value sent to
	// the channel, and done after the channel is closed
	for myVar := range ch {
		fmt.Println("Got value", myVar)
	}

	fmt.Println("All done")
}

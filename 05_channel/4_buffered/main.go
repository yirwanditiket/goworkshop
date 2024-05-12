package main

import (
	"fmt"
)

func main() {
	// creating channel to send and receive "int"
	// buffered with buffer size: 2, bi-directional
	ch := make(chan int, 2)

	// sending to buffered channel will not block if there
	// is space in the buffer
	ch <- 2
	ch <- 3

	// receiving from buffered channel works the same way
	// with unbuffered channel
	myVar := <-ch

	fmt.Println(myVar)
}

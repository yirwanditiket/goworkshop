package main

import (
	"fmt"
)

func main() {
	// creating channel to send and receive "int"
	// unbuffered, bi-directional
	ch := make(chan int)

	// to send data into a channel use ch <- value
	// sending to unbuffered channel blocks if there is no
	// one receiving from the channel
	ch <- 2

	// to receive data from a channel use <- ch
	// receiving from unbuffered channel blocks if there is no
	// one sending to the channel
	myVar := <-ch

	fmt.Println(myVar)
}

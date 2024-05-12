package main

import (
	"fmt"
)

func main() {
	// creating channel to send and receive "int"
	// unbuffered, bi-directional
	ch := make(chan int)

	go func() {
		ch <- 2
	}()

	// myVar will be "int"
	myVar := <-ch

	// no need waitgroup or sleep
	// the channel also act as syncrhonization
	// to wait for the goroutine to finish
	fmt.Println(myVar)
}

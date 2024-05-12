package main

import (
	"fmt"
)

func main() {
	// creating channel to send and receive "int"
	// unbuffered, bi-directional
	ch := make(chan int)

	go func() {
		for i := 0; i <= 2; i++ {
			ch <- i
		}
	}()

	// myVar - myVar3 will be "int"
	myVar := <-ch
	myVar2 := <-ch
	myVar3 := <-ch

	fmt.Println(myVar)
	fmt.Println(myVar2)
	fmt.Println(myVar3)
}

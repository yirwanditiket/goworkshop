package main

import (
	"fmt"
	"sync"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer func() {
			wg.Done()
			fmt.Println("Select exit")
		}()

		select {
		case intCh <- 10:
		case strCh <- "hello":
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("My int", <-intCh)
	}()

	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("My string", <-strCh)
	// }()

	wg.Wait()
	fmt.Println("All done")
}

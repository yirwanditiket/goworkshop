package main

import (
	"fmt"
	"sync"
)

func main() {
	main1()
}

func main1() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("world")
	}()

	wg.Wait()
	fmt.Println("hello")
}

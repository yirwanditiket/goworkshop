package main

import (
	"sync"
)

func main() {
	asc()
}

func asc() {
	x := 20

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		x++
	}()

	wg.Wait()
	println(x)
}

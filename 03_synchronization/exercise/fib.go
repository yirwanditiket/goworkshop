package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()

	total := seq(input)

	fmt.Println("Total:", total)
	fmt.Println("Elapsed:", time.Since(st))
}

func seq(in []int) int {
	total := 0
	for _, i := range in {
		total += fib(i)
		fmt.Println("Done", i)
	}
	return total
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

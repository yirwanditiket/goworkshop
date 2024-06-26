package main

import (
	"fmt"
	"time"
)

func main() {
	main2()
}

func main1() {
	go func() {
		fmt.Println("world")
	}()

	fmt.Println("hello")
}

func main2() {
	go func() {
		fmt.Println("world")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("hello")
}

func main3() {
	go func() {
		fmt.Println("world")
	}()

	fmt.Println("hello")

	time.Sleep(1 * time.Second)
}

func blocking() {
	fmt.Println("1 + 1 = ", 1 + 1)
}

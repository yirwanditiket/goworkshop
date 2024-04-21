package main

import (
	"fmt"
	"time"
)

func main() {
	total := 0

	for i := 0; i < 10; i++ {
		go func() {
			temp := total + 1
			fmt.Println("add")
			total = temp
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(total)
}

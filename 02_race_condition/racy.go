package main

import (
	"fmt"
	"time"
)

func main() {
	total := 0

	for i := 0; i < 10; i++ {
		go func() {
			total = total + 1
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println(total)
}

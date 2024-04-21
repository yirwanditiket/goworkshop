package main

import (
	"fmt"
)

func main() {
	total := 0

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("add -- #", i)
			total = total + 1
		}(i)
	}

	fmt.Println(total)
}

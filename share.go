package main

import (
	"fmt"
	"os"
)

const goroutines = 10

func main() {
	counter := make(chan int, 1)

	for i := 0; i < goroutines; i++ {
		go func(counter chan int) {
			value := <-counter
			value++
			fmt.Println("counter", value)
			if value == goroutines {
				os.Exit(0)
			}
			counter <- value
		}(counter)
	}

	counter <- 0
	for {

	}
}

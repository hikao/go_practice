package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutine = 3

func main() {
	c := make(chan int)

	for i := 0; i < goroutine; i++ {
		go func(s chan<- int) {
			time.Sleep(
				time.Duration(rand.Int63n(10)) * time.Second)
			fmt.Println("completed")
			s <- 0
		}(c)
	}

	for i := 0; i < goroutine; i++ {
		<-c
	}

	fmt.Println("all completed")
}

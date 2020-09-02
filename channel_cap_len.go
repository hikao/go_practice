package main

import "fmt"

func main() {
	// make channel
	c := make(chan int, 10)

	c <- 1
	c <- 10
	fmt.Println("cap:", cap(c))
	fmt.Println("len:", len(c))

	val := <-c
	fmt.Println("val:", val)
	fmt.Println("cap:", cap(c))
	fmt.Println("len:", len(c))
}

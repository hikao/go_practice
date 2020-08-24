package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("10")
}

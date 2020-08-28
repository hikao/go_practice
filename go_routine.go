package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	fmt.Println("call test")
	test()

	fmt.Println("call go routine")
	go test()
	time.Sleep(3 * time.Second)

	fmt.Println("finished")
}

func test() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	time.Sleep(1 * time.Second)
}

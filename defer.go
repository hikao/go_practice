package main

import "fmt"

func main() {
	fmt.Println("start")

	defer delay()

	fmt.Println("end")

}

func delay() {
	fmt.Println("called delay func")
}

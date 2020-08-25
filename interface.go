package main

import "fmt"

type Calculator interface {
	Calculate(a int, b int) int
}

type Add struct {
}

func (x Add) Calculate(a int, b int) int {
	return a + b
}

type Sub struct {
}

func (y Sub) Calculate(a int, b int) int {
	return a - b
}

func main() {
	var add Add
	var sub Sub

	var cal Calculator

	cal = add

	fmt.Println("add:", cal.Calculate(8, 3))

	cal = sub

	fmt.Println("sub:", cal.Calculate(8, 3))

	var empty interface{}
	empty = "abc"
	fmt.Println(empty)
	empty = 123
	fmt.Println(empty)
	empty = 2.56
	fmt.Println(empty)

}

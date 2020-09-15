package main

import "fmt"

type hogeType int

func (h hogeType) setByValue(newValue hogeType) {
	h = newValue
}

func (h *hogeType) setByPointer(newValue hogeType) {
	*h = newValue
}

func (h *hogeType) add(increment hogeType) hogeType {
	*h += increment
	return *h
}

func main() {
	var x hogeType = 0

	x.setByValue(10)
	fmt.Println(x)

	x.setByPointer(20)
	fmt.Println(x)

	fmt.Println(x.add(30))

	a := x.add

	fmt.Println(a(50))
}

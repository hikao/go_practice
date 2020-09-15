package main

import "fmt"

type hogeType int

func (h hogeType) println() {
	fmt.Println(h)
}

func main() {
	var i hogeType = 222
	i.println()
}

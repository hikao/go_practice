package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p1 Person
	p1.name = "Jhon"
	p1.age = 10

	p2 := Person{age: 31, name: "Bob"}

	p3 := Person{"Tom", 49}

	p4 := &Person{"Mike", 40}
	p4.name = "A"

	fmt.Println(p1, p2, p3, *p4)
}

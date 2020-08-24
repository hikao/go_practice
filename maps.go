package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	fmt.Println(m)
	m["Bell Labs"] = Vertex{
		40.21111, -74.23222,
	}
	fmt.Println(m["Bell Labs"])
}

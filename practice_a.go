package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	var a, b, c int
	var s string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)
	fmt.Println(a+b+c, s)
}

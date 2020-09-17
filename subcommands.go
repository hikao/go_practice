package main

import (
	"flag"
	"fmt"
)

func main() {
	help := flag.String("a", "all", "-a needs string")
	flag.Parse()
	fmt.Println(*help)
}

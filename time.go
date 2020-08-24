package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Add(-24 * time.Hour)
	a := t.Format(time.RFC3339)
	fmt.Println(a)
}

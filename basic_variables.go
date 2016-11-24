package main

import (
  "fmt"
  "strings"
)

func main() {
    var s string
    var i int
    var f float64
    var b bool
    fmt.Println(s)
    fmt.Println(i)
    fmt.Println(f)
    fmt.Println(b)
    fmt.Println(nil)

    //Printf
    fmt.Println(strings.Repeat("*", 10))
    fmt.Printf("Printf\n")
    fmt.Printf("s: %s, i: %d, f: %f, b: %t", s, i, f, b)
}

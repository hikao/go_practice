package main

import "fmt"

func main() {
  const (
    sun = iota
    man
    tue
    wed
    thu
    fri
    sat
  )
  fmt.Println(sun, man, tue, wed, thu, fri, sat)
}


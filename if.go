package main

import "fmt"

func main(){
  a :=1
  if a < 0 {
    fmt.Println("a is negative")
  } else if (a == 0) {
    fmt.Println("a is 0")
  } else {
    fmt.Println("a is positive")
  }
}

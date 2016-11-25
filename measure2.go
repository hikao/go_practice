package main

import (
  "log"
  "time"
)

func main() {
  log.Print("started")

  c1 := make(chan bool)
  c2 := make(chan bool)
  c3 := make(chan bool)

  go func() {
    log.Print("hoge")
    time.Sleep(3 * time.Second)
    c1 <- true
  }()

  go func(){
    log.Print("fuga")
    time.Sleep(3 * time.Second)
    c2 <- true
  }()

  go func(){
    log.Print("moge")
    time.Sleep(3 * time.Second)
    c3 <- true
  }()

  <- c1
  <- c2
  <- c3
  log.Print("Finish")
}

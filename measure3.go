package main

import (
  "log"
  "time"
)

func main() {
  log.Print("started")

  c := make(chan bool)

  funcs := []func(){
    func() {
      log.Print("hoge")
      time.Sleep(3 * time.Second)
      c <- true
    },
    func(){
      log.Print("fuga")
      time.Sleep(3 * time.Second)
      c <- true
    },
    func(){
      log.Print("moge")
      time.Sleep(3 * time.Second)
      c <- true
    },
  }
  for _, sleep := range funcs {
    go sleep ()
  }

  for i := 1; i <= len(funcs); i++ {
    <- c
  }
  log.Print("Finish")
}

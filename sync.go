package main

import (
  "log"
  "time"
  "sync"
)

func main() {
  log.Print("started")

  funcs := []func(){
    func() {
      log.Print("hoge")
      time.Sleep(3 * time.Second)
    },
    func(){
      log.Print("fuga")
      time.Sleep(3 * time.Second)
    },
    func(){
      log.Print("moge")
      time.Sleep(3 * time.Second)
    },
  }

  var waitGroup sync.WaitGroup
  for _, sleep := range funcs {
    waitGroup.Add(1)
      go func(function func()){
        defer waitGroup.Done()
        function()
      }(sleep)
  }

  waitGroup.Wait()
  log.Print("Finish")
}


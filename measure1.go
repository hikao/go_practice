package main

import (
  "log"
  "time"
)

func main() {
  log.Print("started")

  log.Print("hoge")
  time.Sleep(3 * time.Second)

  log.Print("fuga")
  time.Sleep(3 * time.Second)

  log.Print("moge")
  time.Sleep(3 * time.Second)

  log.Print("Finished")
}

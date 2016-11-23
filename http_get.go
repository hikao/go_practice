package main

import (
  "net/http"
  "fmt"
  "io"
)

func main() {
  res, err := http.Get("http://google.co.jp/")
  if err != nil {
    fmt.Println("Request error", err)
    return
  }
  defer res.Body.Close()

  buf := make([]byte, 256)
  for {
    n, err := res.Body.Read(buf)
    if n == 0 || err == io.EOF {
      break;
    } else if err != nil {
      fmt.Println("Read response body error:", err)
      return
    }
    fmt.Println(string(buf[:n]))
  }
}

package main

import (
  "io/ioutil"
  "fmt"
  "gopkg.in/yaml.v2"
)

func main() {
  buf, err := ioutil.ReadFile("status_checker/urls.yml")
  if err != nil {
    panic(err)
  }

  d := make(map[interface{}]interface{})
  err = yaml.Unmarshal(buf, &d)
  fmt.Printf("%s\n", d["urls"])
}

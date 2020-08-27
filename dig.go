package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("dig", "example.com").Output()
	if err != nil {
		log.Printf("Failed to run command:%v", err)
	}
	fmt.Println(string(out))
}

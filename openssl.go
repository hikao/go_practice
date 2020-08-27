package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	flag.Parse()
	args := flag.Args()
	cmd := fmt.Sprintf("openssl s_client -connect %v:443 | openssl x509 -noout -enddate", args[0])
	fmt.Println(cmd)
	o, _ := exec.Command("sh", "-c", cmd).Output()
	fmt.Println(string(o))
}

package main

import (
    "fmt"
	"io"
    "net/http"
)

func main() {
	res, err := http.Get("https://golang.org")
	if err != nil {
		fmt.Println("Request error:", err)
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

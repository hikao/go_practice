package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "This is a test of strings in golang."

	// slice
	fmt.Println(s[:5])
	fmt.Println(s[5:7])
	fmt.Println(s[5:])
	fmt.Println(s[:])

	space := "  aaa aaa  aaaaa   "
	fmt.Println(strings.TrimSpace(space))

	csv := "a,b,c,d,e,f,g"
	fmt.Println(strings.Split(csv, ","))

	// str to int
	i64, err := strconv.ParseInt("123456789", 10, 0)
	fmt.Println(int(i64), err)

	// int to str
	var i int64 = 123456789

	hoge := strconv.FormatInt(i, 10)
	fmt.Println(hoge)
	fmt.Printf("%T\n", hoge)

}

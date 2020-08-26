package main

import "log"

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func main() {
	err := do()

	if err == nil {
		log.Println("no error")
	} else {
		log.Println("err???")
	}
}

func do() error {
	//var ret *MyError = nil
	//	return ret
	return nil
}

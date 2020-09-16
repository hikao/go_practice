package main

import "fmt"

type User struct {
	Age string
}

func (u User) setAgeByValue(age string) {
	fmt.Printf("address is %p\n", &u)
	u.Age = age
}

func (u *User) setAgeByPtr(age string) {
	fmt.Printf("address is %p\n", u)
	u.Age = age
}

func main() {
	user := new(User)
	user.setAgeByValue("30")
	fmt.Println(user.Age)

	user.setAgeByPtr("25")
	fmt.Println(user.Age)

	user.setAgeByValue("30")
	fmt.Println(user.Age)

	user.setAgeByPtr("25")
	fmt.Println(user.Age)
}

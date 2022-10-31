package main

import (
	"fmt"
)

type User struct {
	Name string
}

func main() {
	// getUser()は *Userを返す。
	user := getUser()
	fmt.Println(getName(user))
}

func getName(user *User) string {
	if user == nil {
		return ""
	}
	return user.Name
}

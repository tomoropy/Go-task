package main

import "fmt"

type User struct {
	Name string
}

var user = &User{"tomo"}

func main() {
	// getUser()は *Userを返す。
	if user := getUser(); user == nil {
		fmt.Println("userを定義してください")
	}else{
		fmt.Println(getName(user))
	}
}

func getName(user *User) string {
	return user.Name
}

func getUser() *User {
	return user
}

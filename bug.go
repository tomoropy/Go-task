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
	// Nameがあるか確認
	if user.Name == "" {
		return "userの名前がありません"
	}
	return user.Name
}

func getUser() *User {
	var user = &User{} //userのNameがない状態で初期化
	return user
}

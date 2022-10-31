// CanAccess()を別の方法で書き換えてください。引数や戻り値は変えずに関数の内部のみ変更すること。
package main

import "fmt"

type Role int

const (
	STUDENT  Role = 1
	MENTOR   Role = 2
	GRADUATE Role = 3
	ADMIN    Role = 4
)

type aUser struct {
	Role
	Name string
}

func (user *aUser) CanAccess() bool {
	return user.Role == MENTOR || user.Role == ADMIN
}

func main() {
	user := aUser{Role: 2, Name: "tarou"}
	fmt.Println(user.CanAccess())
}

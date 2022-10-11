// このコードの実行結果はどうなるか予想してください。
// 実際に実行してなぜこのような挙動になるのかを説明してください。 また、正しいと思われる挙動にするための修正をしてください。

package	user 

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{"tarou", 33},
		{"zirou", 22},
		{"itirou", 11},
	}

	// ✖︎ for _, user := range users {
	for i, _ := range users {
		// ✖︎ user.Age = 44
		users[i].Age = 44
	}

	fmt.Printf("%v", users) // どうなる？ -> A.変更されない

	// A.
	// rangeでuserを変数として定義しているので、参照渡しではなく値渡しになり
	// 元の変数userは変わらない
	// インデックスだけもらって中身を書き換えると正常に動作する
}

// 課題[2]
// mapのvalueからkeyを取得する問題。
// 引数にmapとvalueを受け取り、対応するkeyがあればkeyを返却する関数の実装
// map[int]stringの型指定

package main

import (
	"errors"
	"fmt"
)

func main()  {
	m := map[int]string {
		1 : "01",
		2 : "02",
		3 : "03",
	}

	key, err := findKeyByValue(m, "01")
	fmt.Printf("key is : %d\n", key)

	key, err = findKeyByValue(m, "04")
	fmt.Printf("key is : %d\n", key)

	if err != nil {
		fmt.Println(err)
	}

}

func findKeyByValue(m map[int]string, v string)(int, error) {
	for key, value := range m {
		if value == v {
			return key, nil
		}
	}
	return 0, errors.New("does not exits")
}

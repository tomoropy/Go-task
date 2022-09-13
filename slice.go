// 課題[1]
// 1つは {1, “2”, 10, “11"}のスライスをループして、
// 一桁の数字なら01のように0をいれて2桁にして表示する問題

package main

import (
	"fmt"
	"strconv"
)

func main() {
	init := []interface{}{ 1, "2", 10, "11"}

	for _, v := range init {
		switch v.(type) {
		case int:
			fillZero(v.(int))
		case string:
			arg, _ := strconv.Atoi(v.(string))
			fillZero(arg)

		}
	}
}

func fillZero(num int) {
	if num > 10 {
		fmt.Println(num)
	}else {
		fmt.Printf("%02d\n", num)
	}
}

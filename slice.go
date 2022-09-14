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
			fmt.Printf("%02d\n", v.(int))
		case string:
			arg, _ := strconv.Atoi(v.(string))
			fmt.Printf("%02d\n", arg)
		}
	}
}

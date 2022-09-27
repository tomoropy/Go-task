// 課題[3]
// type MyIntSlice []intでIntスライスと互換のある独自型を作り、重複を排除するメソッドを実装する。

// type MyIntSlice []int
// m := MyIntSlice{1, 2, 2, 3, 3, 3, 4, 5}
// fmt.Println(m.Unique()) // [1, 2, 3, 4, 5]

package main

import "fmt"


type MyIntSlice []int

func main()  {
	m := MyIntSlice{1, 2, 2, 3, 3, 3, 4, 5}

	fmt.Println(m.Unique())
}

func (m MyIntSlice) Unique() (unique []int) {
	validMap := make(map[int]bool)
	unique = []int{}

	for _, v := range m{
		if !validMap[v] {
			validMap[v] = true
			unique = append(unique, v)
		}
	}
	return
}

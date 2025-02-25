package main

import "fmt"

/*
格式:
	for key, val := range coll {
    ...
}
*/

func main() {
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	str := "ms的go教程"
	// 因为一个字符串是 Unicode 编码的字符（或称之为 rune ）集合
	// char 实际类型是 rune 类型
	for pos, char := range str {
		fmt.Println(pos, char)
	}
}

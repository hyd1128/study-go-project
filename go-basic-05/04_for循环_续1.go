package main

import "fmt"

func main() {
	s := "abcd"
	// 初始
	// // 这样写多次调用length函数
	// for i := 0; i < length(s); i++ {
	// 	println(i, s[i])
	// }

	// 优化一波
	for i, n := 0, length(s); i < n; i++ {
		fmt.Println(i, s[i])
	}
}

func length(s string) int {
	println("call length.")
	return len(s)
}

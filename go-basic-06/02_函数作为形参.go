package main

import "fmt"

func testOne(fn func() int) int {
	return fn()
}
func fn() int {
	return 200
}
func main() {
	// 这是直接使用匿名函数
	// s1 := testOne(func() int { return 100 })
	// 这是传入一个函数
	s1 := testOne(fn)
	fmt.Println(s1)
}

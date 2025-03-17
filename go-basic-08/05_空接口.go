package main

import "fmt"

/*
什么是空接口:
1. 空接口是指没有定义任务方法的接口
2. 任何类型都可以实现空接口
3. 空接口类型的变量可以储存任意类型的变量
*/
func main() {
	// 定义一个空接口
	var x interface{}
	s := "ms的go语言教程"
	x = s
	fmt.Printf("type: %T value: %v\n", x, x)

	i := 100
	x = i
	fmt.Printf("type: %T value: %v\n", x, x)

	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}

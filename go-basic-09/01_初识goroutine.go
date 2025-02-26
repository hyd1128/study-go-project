package main

import (
	"fmt"
	"time"
)

/*
实现一个goroutine的格式:
	go 函数名(参数列表)
	- 函数名: 要调用的函数名
	- 参数列表: 调用红函数需要传入的参数
*/

// 启动单个goroutine
func hello(num int) {
	fmt.Println("Hello Goroutine!", num)
}

// func main() {
// 	hello()
// 	fmt.Println("main goroutine done!")
// }

// 使用goroutine运行hello
// func main() {
// 	go hello()
// 	fmt.Println("main goroutine done!")
// 	time.Sleep(time.Second)
// 	// 由于在创建新的goroutine的时候需要花费时间, 所以新创建的额goroutine要晚于主goroutine执行
// }

// 启动多个goroutine
func main() {
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second * 2)
}

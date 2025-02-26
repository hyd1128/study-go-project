package main

import "fmt"

/*
声明channel的格式:
	var 变量 chan 类型
*/

// func main() {
// 	ch := make(chan int)
// 	ch <- 10
// 	fmt.Println("发送成功")
// }
// 上述抱错: fatal error: all goroutines are asleep - deadlock!

// 解决办法
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

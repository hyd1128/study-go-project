package main

import "fmt"

/*
空接口的应用
*/

// 1. 空接口作为函数的参数
// 使用空接口实现可以接收任意类型的函数参数
func show(a interface{}) {
	fmt.Printf("type: %T value: %v \n", a, a)
}

func main() {
	// 2. 空接口作为map的值
	// 使用空接口实现可以保存任意值的字典
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "李白"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

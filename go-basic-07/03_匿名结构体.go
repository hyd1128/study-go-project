package main

import "fmt"

// 打印消息 传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Println("%T, msg: %v", msg, msg)
}

func main() {
	/*
		格式:
				ins := struct {
			    // 匿名结构体字段定义
			    字段1 字段类型1
			    字段2 字段类型2
			    …
			}{
			    // 字段值初始化
			    初始化字段1: 字段1的值,
			    初始化字段2: 字段2的值,
			    …
			}
	*/

	/*
		& 和 * 的配合使用
		& 是取地址操作符，返回某个变量的内存地址。
		* 是解引用操作符，通过地址取值。

		在变量前：*p 表示解引用，获取指针指向的值。
		在类型前：*int 表示这是一个指向 int 的指针类型。

		* 获取值：当用在指针变量前，表示解引用，获取地址中的值。
		* 表示地址：在类型声明中，表示这是一个指针类型，变量存储的是地址。
	*/
	// 实例化一个匿名结构体
	msg := &struct {
		id   int
		data string
	}{
		// 初始化部分
		1024, "hello",
	}
	printMsgType(msg)

}

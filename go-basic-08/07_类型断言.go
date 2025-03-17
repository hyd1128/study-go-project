package main

import (
	"fmt"
)

/*
接口值
类型断言的格式
x.(T)

x: 表示类型为interface{}的变量
y: 表示断言x可能是的类型

该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
*/

func justifyType(x interface{}) {
	// x.(type) 不是取 x 的值，而是获取 x 实际的动态类型。
	// switch x.(type) 语法只能用于 type switch，它不是普通的 switch 语句
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}

}

func main() {
	// var x interface{}
	// x = "ms的go教程"
	// v, ok := x.(string)
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("类型断言失败")
	// }

	// 断言多次
	justifyType("123")

}

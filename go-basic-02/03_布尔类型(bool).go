package main

import "fmt"

// 定义布尔类型变量
// var 变量名 bool

// `==`,`>`,`<`，`<=`, `>=`,`&&(AND)`,`||(OR)`等都会产生bool值
func main() {
	// var aVar = 10
	// fmt.Println(aVar == 5)
	// fmt.Println(aVar == 10)
	// fmt.Println(aVar != 5)
	// fmt.Println(aVar != 10)

	/*
			Go语言对于值之间的比较有非常严格的限制，只有两个相同类型的值才可以进行比较，
			如果值的类型是接口（interface），那么它们也必须都实现了相同的接口。

			如果其中一个值是`常量`，那么另外一个值可以不是常量，但是类型必须和该常量类型相同。

			如果以上条件都不满足，则必须将其中一个值的类型转换为和另外一个值的类型相同之后才可以进行比较。

		    `&&(AND)`,`||(OR)`是具有短路行为的，如果运算符左边的值已经可以确定整个布尔表达式的值，
			那么运算符右边的值将不再被求值。(&&优先级高于||)
	*/

	var a = 10
	if a > 11 && a < 30 {
		fmt.Println("正确")
	} else {
		fmt.Println("错误")
	}

	if a > 5 || a < 30 {
		fmt.Println("正确")
	} else {
		fmt.Println("错误")
	}

	// 布尔型数据只有true和false，且不能参与任何计算以及类型转换
}


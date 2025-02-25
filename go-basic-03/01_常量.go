package main

import "fmt"

func main() {
	// const pi = 3.14159
	// fmt.Println(pi)

	// 批量声明变量
	const (
		e  = 2.7182818
		pi = 3.14159
	)

	fmt.Println(e)
	fmt.Println(pi)

	/*
		如果是批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略，
		如果省略初始化表达式则表示使用前面常量的初始化表达式，对应的常量类型也是一样的
	*/
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)

	// iota常量生成器
	const (
		sunday = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)

	/*
		在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加1
	*/
	fmt.Println(sunday,
		monday,
		tuesday,
		wednesday,
		thursday,
		friday,
		saturday)
}

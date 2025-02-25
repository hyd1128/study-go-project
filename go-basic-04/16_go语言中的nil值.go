package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// fmt.Println(nil == nil) // invalid operation: nil == nil (operator == not defined on untyped nil)
	/*
			在Go语言中，布尔类型的零值（初始值）为 false，数值类型的零值为 0，
			字符串类型的零值为空字符串`""`，而指针、切片、映射、通道、函数和接口的零值则是 nil。

			nil和其他语言的null是不同的。

			nil 标识符是不能比较的

		nil值特点:
			1. nil 标识符是不能比较的
			2. nil 不是关键字或保留字
			3. nil 并不是Go语言的关键字或者保留字，也就是说我们可以定义一个名称为 nil 的变量
			4. nil 没有默认类型
			5.不同类型 nil 的指针是一样的
	*/

	// error :use of untyped nil
	// fmt.Printf("%T", nil)
	// print(nil)

	// var arr []int
	// var num *int
	// fmt.Printf("%p\n", arr)
	// fmt.Printf("%p", num)

	// nil 是 map、slice、pointer、channel、func、interface 的零值
	// var m map[int]string
	// var ptr *int
	// var c chan int
	// var sl []int
	// var f func()
	// var i interface{}
	// fmt.Printf("%#v\n", m)
	// fmt.Printf("%#v\n", ptr)
	// fmt.Printf("%#v\n", c)
	// fmt.Printf("%#v\n", sl)
	// fmt.Printf("%#v\n", f)
	// fmt.Printf("%#v\n", i)

	/*
		零值是Go语言中变量在声明之后但是未初始化被赋予的该类型的一个默认值。

		不同类型的 nil 值占用的内存大小可能是不一样的
	*/
	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8
	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24
	var m map[int]bool
	fmt.Println(unsafe.Sizeof(m)) // 8
	var c chan string
	fmt.Println(unsafe.Sizeof(c)) // 8
	var f func()
	fmt.Println(unsafe.Sizeof(f)) // 8
	var i interface{}
	fmt.Println(unsafe.Sizeof(i)) // 16


}

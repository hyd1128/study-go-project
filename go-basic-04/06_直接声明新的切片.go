package main

import "fmt"

func main() {
	/*
		name 表示切片的变量名，Type 表示切片对应的元素类型。
		var name []Type
	*/
	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	fmt.Println(strList)
	fmt.Println(numList)
	fmt.Println(numListEmpty)
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
	/*
		切片是动态结构，只能与 nil 判定相等，不能互相判定相等。声明新的切片后，可以使用 append() 函数向切片中添加元素。
	*/
	var strListOne []string
	// 追加一个元素
	strListOne = append(strListOne, "go教程")
	fmt.Println(strListOne)

}

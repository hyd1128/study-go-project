package main

import (
	"fmt"
)

func main() {
	/*
		array_name 为数组的名字，array_type 为数组的类型，size1、size2 等等为数组每一维度的长度。
		var array_name [size1][size2]...[sizen] array_type
	*/
	// 声明一个二维数组
	var arr1 [4][2]int
	fmt.Println(arr1)

	// 使用数组字面量声明并初始化一个二维整形数组
	arr2 := [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(arr2)

	// 声明并初始化数组中索引为1和3的元素
	arr3 := [4][2]int{1: {1, 2}, 3: {3, 4}}
	fmt.Println(arr3)

	// 取值
	fmt.Println(arr3[1][0])

	// 循环取值
	for index, value := range arr3 {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
}

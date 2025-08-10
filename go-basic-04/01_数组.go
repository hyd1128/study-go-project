package main

import "fmt"

func main() {
	/*
		数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。
		因为数组的长度是固定的，所以在Go语言中很少直接使用数组。
		var 数组变量名 [元素数量]Type
	*/
	var arr [3]int // 赋予默认值
	fmt.Println(arr)

	// 通过数组下标取值
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	// for range取值
	for index, value := range arr {
		fmt.Printf("索引: %d, 值: %d \n", index, value)
	}
}

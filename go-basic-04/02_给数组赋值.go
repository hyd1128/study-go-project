package main

import (
	"fmt"
)

func main() {
	// 初始化的时候赋值
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// 如果第三个赋值,则为默认值
	var arr1 [3]int = [3]int{1, 2}
	fmt.Println(arr1)

	// 简单声明
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// 不写数据数量, 而使用..., 表示数组的长度是更具初始化值的个数来计算的
	arr3 := [...]int{1, 3, 4}
	fmt.Println(arr3)
	// 一定要注意，数组是定长的，不可更改，在编译阶段就决定了

	// 只想给第三个值赋值,其余的值为默认值
	arr4 := [3]int{2: 3}
	for index, value := range arr4 {
		fmt.Printf("索引值: %d, 值: %d", index, value)
	}
}

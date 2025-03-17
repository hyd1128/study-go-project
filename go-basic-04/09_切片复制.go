package main

import "fmt"

func main() {
	/*
		copy( destSlice, srcSlice []T) int
	*/

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	// copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	// fmt.Println(result1)
	// fmt.Println(result2)
	fmt.Println(slice1)
	fmt.Println(slice2)
	/*
		Go语言的内置函数 copy() 可以将一个数组切片复制到另一个数组切片中，如果加入的两个数组切片不一样大，
		就会按照其中较小的那个数组切片的元素个数进行复制。

		其中 `srcSlice `为数据来源切片，`destSlice `为复制的目标（也就是将 srcSlice 复制到 destSlice），
		`目标切片必须分配过空间且足够承载复制的元素个数`，并且来源和目标的`类型必须一致`，copy() 函数的返回值表示实际发生复制的元素个数。
	*/

}

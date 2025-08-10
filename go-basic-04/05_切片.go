package main

import (
	"fmt"
)

func main() {
	// 切片语法
	/*
		slice [开始位置 : 结束位置]
			语法说明如下：
			- slice：表示目标切片对象；
			- 开始位置：对应目标切片对象的索引；
			- 结束位置：对应目标切片的结束索引。
	*/
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:3])
	/*
			从数组或切片生成新的切片拥有如下特性：

		- 取出的元素数量为：结束位置 - 开始位置；
		- 取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
		- 当缺省开始位置时，表示从连续区域开头到结束位置`(a[:2])`；
		- 当缺省结束位置时，表示从开始位置到整个连续区域末尾`(a[0:])`；
		- 两者同时缺省时，与切片本身等效`(a[:])`；
		- 两者同时为 0 时，等效于空切片，一般用于切片复位`(a[0:0])`。
	*/
	var highRaiseBUilding [30]int
	// 开始条件 结束条件 过程操作
	for i := 0; i < 30; i++ {
		highRaiseBUilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRaiseBUilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRaiseBUilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRaiseBUilding[:10])

	slice_one := highRaiseBUilding[10:15]
	fmt.Println(slice_one)
	fmt.Printf("数组类型: %T \n", highRaiseBUilding)
	fmt.Printf("切片类型: %T \n", slice_one)
	fmt.Println(slice_one[len(slice_one)-1])
}

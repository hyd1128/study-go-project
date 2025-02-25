package main

import (
	"fmt"
)

// 匿名函数定义格式

/*
func (参数列表) (返回参数列表) {
		函数体
	}
*/

// 匿名函数用作回调函数
// 遍历切片的每一个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {
	// // 这里将一个函数当作一个变量一样操作
	// getSqrt := func(a float64) float64 {
	// 	// 接收一个float64的数字, 并返回这个数字的平方根
	// 	return math.Sqrt(a)
	// }
	//
	// fmt.Println(getSqrt(4))

	// // 直接调用匿名函数
	// func(data int) {
	// 	fmt.Println("hello", data)
	// }(100)

	// 使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})

}

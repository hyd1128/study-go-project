package main

import "fmt"

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6]
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	// 这打印出来长度为2
	fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, len(myslice))
	myslice = myslice[:cap(myslice)]
	// 为什么 myslice 的长度为2，却能访问到第四个元素
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
	fmt.Println("\n", len(myslice))

	// 注意: 数组的切片开始切片索引既决定长度同时也决定容量，而结束切片索引只决定长度，不决定容量
	// 计算公式: 切片数组的长度: 结束索引 - 开始索引
	// 切片数组的容量: 原数组容量 - 开始索引
}

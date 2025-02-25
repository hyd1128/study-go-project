package main

import "fmt"

func main() {
	// 定义数组
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Printf("arr类型%T, arr:", arr1, arr1)
	fmt.Println()
	arr2 := [3]int{1, 2, 3}
	fmt.Printf("arr类型%T, arr:", arr2, arr2)
	fmt.Println()
	arr3 := [...]int{1, 2, 3}
	fmt.Printf("arr类型%T, arr:", arr3, arr3)
	fmt.Println()
	// 定义切片
	slice1 := arr1[0:len(arr1)]
	fmt.Printf("slice类型%T, slice:", slice1, slice1)
	fmt.Println()
	var slice2 []int
	fmt.Printf("slice类型%T, slice:", slice2, slice2)
	fmt.Println()
	var slice3 = []int{}
	fmt.Printf("slice类型%T, slice:", slice3, slice3)
	fmt.Println()
	slice4 := []int{}
	fmt.Printf("slice类型%T, slice:", slice4, slice4)
	fmt.Println()
	slice5 := make([]int, 5, 10)
	fmt.Printf("slice类型%T, slice:", slice5, slice5)
	fmt.Println()
	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr4[:])
	fmt.Println(arr4[:len(arr4)])

	// :=和copy在数组和切片上的区别
	arr8 := [5]int{1, 2, 3, 4, 5}
	// arr9 := arr8
	// var arr10 [5]int
	// copy(arr10, arr8) // copy()方法不能用于数组赋值，只能够用于切片复制
	// fmt.Println(arr8)
	// fmt.Println(arr9) // 通过实验可知，数组赋值是深度拷贝
	slice10 := arr8[:]
	// fmt.Println(arr8)
	slice11 := make([]int, 5)
	copy(slice11, arr8[:])
	arr8[0] = 10000
	fmt.Println(arr8)
	fmt.Println(slice10) // 通过实验可知，切片赋值是浅拷贝
	fmt.Println(slice11)

}

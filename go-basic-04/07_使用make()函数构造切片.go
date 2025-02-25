package main

import "fmt"

func main() {
	// 如果需要动态地创建一个切片，可以使用 make() 内建函数，格式如下：
	// make( []Type, size, cap )
	// Type` 是指切片的元素类型，`size` 指的是为这个类型分配多少个元素，`cap `为预分配的元素数量，
	// 这个值设定后不影响 size，只是能提前分配空间，`降低多次分配空间造成的性能问题`。

	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))

}

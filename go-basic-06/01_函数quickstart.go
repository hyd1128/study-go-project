package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(1, 10))
	fmt.Println(max(-1, -2))

	fmt.Println(test(1, 2, "123"))

	x, y := testTwo(1, 2, "abc")
	fmt.Println(x, y)
}

// 类型相同的相邻参数，参数类型可合并。
func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

/*
**Go语言是编译型语言，所以函数编写的顺序是无关紧要的，鉴于可读性的需求，
最好把 main() 函数写在文件的前面，其他函数按照一定逻辑顺序进行编写（例如函数被调用的顺序）。**
*/

func test(x, y int, s string) (int, string) {
	// 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func testTwo(x, y int, z string) (int, string) {
	resultX := x + y
	resultY := z
	return resultX, resultY
}

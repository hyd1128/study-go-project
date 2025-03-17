package main

import "fmt"

// 定义点结构
type PointOne struct {
	X int
	Y int
}

// 非指针接收器的加方法
func (p PointOne) Add(other PointOne) PointOne {
	// 成员值与参数相加后返回新的结构
	return PointOne{p.X + other.X, p.Y + other.Y}
}
func main() {
	// 初始化点
	p1 := PointOne{1, 1}
	p2 := PointOne{2, 2}
	// 与另外一个点相加
	result := p1.Add(p2)
	// 输出结果
	fmt.Println(result)
}

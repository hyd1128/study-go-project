package main

import (
	"fmt"
	"strings"
)

func main() {
	// 查找
	tracer := "张三来了,张三bye bye"

	// 正向搜索字符串
	comma := strings.Index(tracer, ",")
	fmt.Println(",所在的位置:", comma)
	fmt.Println(tracer[comma+1:]) // 张三bye bye

	add := strings.Index(tracer, "+")
	fmt.Println("+所在的位置:", add) // +所在的位置: -1

	pos := strings.Index(tracer[comma:], "张三")
	fmt.Println("张三 所在的位置", pos) // 张三，所在的位置 1

	fmt.Println(comma, pos, tracer[5+pos:]) // 12 1 张三bye bye
}

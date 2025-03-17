package main

import (
	"bytes"
	"fmt"
)

func main() {
	str1 := "你好,"
	str2 := "ms的go教程"
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(str1)
	stringBuilder.WriteString(str2)

	// Sprint以字符串形式返回
	result := fmt.Sprint(stringBuilder.String())
	fmt.Println(result)

	/*
		    %c  单一字符
			%T  动态类型
			%v  本来值的输出
			%+v 字段名+值打印
			%d  十进制打印数字
			%p  指针，十六进制
			%f  浮点数
			%b 二进制
			%s string
	*/
}

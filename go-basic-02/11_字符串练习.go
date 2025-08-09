package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello, ms的go教程Java教程"
	source := "Java"
	target := "Go"
	// 找到target字符串的起始位置
	index := strings.Index(str1, source)
	// 判断target字符串的长度
	sourceLength := len(source)
	// 截取target字符串之前的string
	start := str1[:index]
	// 截取target字符串之后的string
	end := str1[index+sourceLength:]

	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(start)
	stringBuilder.WriteString(target)
	stringBuilder.WriteString(end)
	fmt.Println(stringBuilder.String())
}

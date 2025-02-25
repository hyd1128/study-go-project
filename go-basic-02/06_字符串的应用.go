package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	// 计算字符串长度
	str3 := "hello"
	str4 := "你好"

	fmt.Println(len(str3))
	fmt.Println(len(str4))
	fmt.Println(utf8.RuneCountInString(str3))
	fmt.Println(utf8.RuneCountInString(str4))

	// 因为编译器会在行尾自动补全分号，所以拼接字符串用的加号“+”必须放在第一行末尾。
	str := "第一部分" +
		"第二部分"
	fmt.Println(str)

	// 可以使用“+=”来对字符串进行拼接
	str1 := "hel" + "lo"
	str1 += " world"
	fmt.Println(str1)

	str5 := "你好,"
	str6 := "ms的go教程"
	var stringBuilder bytes.Buffer
	// 节省内存分配，提高处理效率
	stringBuilder.WriteString(str5)
	stringBuilder.WriteString(str6)
	fmt.Println(stringBuilder.String())

	// 从含有中文的字符串中提取中文字符
	var myStrOne string = "hello,ms的go教程"
	fmt.Println(string([]rune(myStrOne)[8]))

	var str7 string = "hello"
	var str8 string = "hello,ms的go教程"
	// 遍历
	for i := 0; i < len(str7); i++ {
		fmt.Printf("ascii: %c %d\n", str1[i], str1[i])
	}
	for _, s := range str7 {
		fmt.Printf("unicode: %c %d\n ", s, s)
	}
	// 中文只能用 for range
	for _, s := range str8 {
		fmt.Printf("unicode: %c %d\n ", s, s)
	}
}

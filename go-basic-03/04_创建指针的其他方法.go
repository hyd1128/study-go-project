package main

import "fmt"

func main() {
	str := new(string)
	*str = "学习go语言"
	fmt.Println(*str)
	fmt.Printf("指针类型: %T, 指针值: %s, 指针指向变量的地址值: %p", str, *str, str)

}

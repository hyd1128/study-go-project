package main

import "fmt"

func main() {
	// `字符串是一种值类型，且值不可变，即创建某个文本后将无法再次修改这个文本的内容。`
	// `当字符为 ASCII 码表上的字符时则占用 1 个字节`
	var myStr string = "hello"
	fmt.Println(myStr)

	var str = "ms的go教程 \n 大法好"
	fmt.Print(str)

	fmt.Println(`\t ms的go教程Go大法好`)  // \t ms的go教程Go大法好
	fmt.Println(`\t ms的go教程
 Go大法好`) // 使用反引号 可以进行字符串换行
	// 反引号一般用在 需要将内容进行原样输出的时候 使用

	var mystr01 string = "hello"
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("myStr01: %s\n", mystr01)
	fmt.Printf("myStr02: %s\n", mystr02)

	// 中文三字节，字母一个字节
	var myStr01 string = "hello,ms的go教程"
	fmt.Printf("mystr01: %d\n", len(myStr01))
}

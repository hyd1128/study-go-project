package main

import "fmt"

// func TestPrint(t *testing.T) {
// 	fmt.Print("我是控制台打印，我不换行 可以自己控制换行")
// }
// func TestPrintln(t *testing.T) {
// 	fmt.Println("我是控制台打印，我换行")
// }
// func TestPrintf(t *testing.T) {
// 	fmt.Printf("我是控制台打印，%s \n", "这是格式化占位符信息,可以自己控制换行")
// }

func main() {
	fmt.Print("我是控制台打印，我不换行 可以自己控制换行")
	fmt.Println("我是控制台打印，我换行")
	fmt.Printf("我是控制台打印，%s \n", "这是格式化占位符信息,可以自己控制换行")

}

package main

import "fmt"

func main() {
	// 利用指针修改值
	var num = 10
	modifyFromPoint(num)
	fmt.Println("未使用指针方法外, num1:", num)

	var num2 = 22
	newModifyFromPoint(&num2)
	fmt.Println("使用指针方法外, num2:", num2)

}

// 方法1 未使用指针修改值
func modifyFromPoint(num int) {
	num = 10000
	fmt.Println("未使用指针, 方法内:", num)
}

// 方法2 使用指针修改值
func newModifyFromPoint(ptr *int) {
	*ptr = 1000 // 修改执政地址指向的值
	fmt.Println("使用指针, 方法内:", *ptr)

}

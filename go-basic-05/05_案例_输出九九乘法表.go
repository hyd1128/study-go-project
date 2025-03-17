package main

import "fmt"

func main() {
	// 遍历， 决定处理第多少行
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d  ", y, x, x*y)
		}
		// 手动换行
		fmt.Println("")
	}
}

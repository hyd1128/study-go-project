package main

import "fmt"

func main() {
	// 	给外层循环取名字
OuterLoop:
	for i := 0; i < 2; i++ {
		// InnerLoop:	// 也可以给内层循环取名字
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop
			}
			fmt.Println(j)
		}
	}
}

/*
continue OuterLoop 使得程序跳过内层循环的剩余部分，并继续执行外层循环。
*/

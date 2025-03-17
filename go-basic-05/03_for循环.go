package main

import "fmt"

func main() {
	// sum := 0

	// 写法1:
	// for {
	// 	sum++
	// 	if sum > 100 {
	// 		// break 跳出循环
	// 		fmt.Printf("%d", sum)
	// 		break
	// 	}
	// }

	// 写法2:
	// i := 0; 赋初值，i<10 循环条件 如果为真就继续执行 ；i++ 后置执行 执行后继续循环
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Printf("sum: %d", sum)

	// 写法3:
	// n := 10
	// for n > 0 {
	// 	n--
	// 	fmt.Println(n)
	// }

	// 写法4:
	// step := 2
	// // 初值可以省略，但是;必须有，但是这样写step的作用域就比较大了，脱离了for循环
	// for ; step > 0; step-- {
	// 	fmt.Println(step)
	// }

	// 进一步简化
	// step := 2
	// for step > 0 {
	// 	step--
	// 	fmt.Println(step)
	// }

	// 结束循环的方式
	// 方式1: return
	// step := 2
	// for step > 0 {
	// 	step--
	// 	fmt.Println(step)
	// 	// 执行一次就结束了
	// 	return
	// }
	// // 不会执行
	// fmt.Println("结束之后的语句....")

	// 方式2: break
	// step := 2
	// for step > 0 {
	// 	step--
	// 	fmt.Println(step)
	// 	// 跳出循环,还会继续执行循环外的语句
	// 	break
	// }
	// // 会执行
	// fmt.Println("结束之后的语句....")

	// 方式3: painc
	// step := 2
	// for step > 0 {
	// 	step--
	// 	fmt.Println(step)
	// 	// 报错了，直接结束
	// 	panic("出错了")
	// }
	// // 不会执行
	// fmt.Println("结束之后的语句....")

	// 方式4:
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return

	// 标签
breakHere:
	fmt.Println("done")
}

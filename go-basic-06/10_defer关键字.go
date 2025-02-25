package main

// func main() {
// 	var whatever = [5]int{1, 2, 3, 4, 5}
//
// 	for i := range whatever {
// 		defer fmt.Println(i)
// 	}
// }

// 延迟调用的其他示例
// func main() {
// 	// 参数求值时机
// 	// i := 1
// 	// defer fmt.Println("defer中的i: ", i)
// 	// i = 2
// 	// fmt.Println("main 中的 i: ", i)
//
// 	// 多个defer的执行顺序
// 	defer fmt.Println("第一个 defer")
// 	defer fmt.Println("第二个 defer")
// 	defer fmt.Println("第三个 defer")
// 	fmt.Println("main 函数执行")
// }

// func main() {
// 	start := time.Now()
// 	log.Printf("开始时间为：%v", start)
// 	defer log.Printf("时间差：%v", time.Since(start)) // Now()此时已经copy进去了
// 	// 不受这3秒睡眠的影响
// 	time.Sleep(3 * time.Second)
//
// 	log.Printf("函数结束")
// }

// func main() {
// 	var whatever = [5]int{1, 2, 3, 4, 5}
// 	for i, _ := range whatever {
// 		// 函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.
// 		defer func() { fmt.Println(i) }()
// 	}
// }

// func main() {
// 	var whatever = [5]int{1, 2, 3, 4, 5}
// 	for i, _ := range whatever {
// 		i := i
// 		defer func() { fmt.Println(i) }()
// 	}
// }

package main

import (
	"time"
)

func testOne(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}

func testTwo(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

// func main() {
// 	// 2个管道
// 	outputOne := make(chan string)
// 	outPutTwo := make(chan string)
// 	// 跑两个协程 写数据
// 	go testOne(outputOne)
// 	go testTwo(outPutTwo)
// 	// 用select监控
// 	select {
// 	case s1 := <-outputOne:
// 		fmt.Println("s1=", s1)
// 	case s2 := <-outPutTwo:
// 		fmt.Println("s2=", s2)
// 	}
// }

// // 当有多个管道同时存在数据, 则随机选择一个管道执行
// func main() {
// 	// 创建2个管道
// 	intChan := make(chan int, 1)
// 	stringChan := make(chan string, 1)
// 	go func() {
// 		// time.Sleep(2 * time.Second)
// 		intChan <- 1
// 	}()
// 	go func() {
// 		stringChan <- "hello"
// 	}()
// 	select {
// 	case value := <-intChan:
// 		fmt.Println("int:", value)
// 	case value := <-stringChan:
// 		fmt.Println("string:", value)
// 	}
// 	fmt.Println("main结束")
// }

// 判断管道有没有存满
// 判断管道有没有存满
// func main() {
// 	// 创建管道
// 	output1 := make(chan string, 10)
// 	// 子协程写数据
// 	go write(output1)
// 	// 取数据
// 	for s := range output1 {
// 		fmt.Println("res:", s)
// 		time.Sleep(time.Second)
// 	}
// }
//
// func write(ch chan string) {
// 	for {
// 		select {
// 		// 写数据
// 		case ch <- "hello":
// 			fmt.Println("write hello")
// 		default:
// 			fmt.Println("channel full")
// 		}
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

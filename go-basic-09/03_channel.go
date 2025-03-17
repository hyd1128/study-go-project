package main

/*
声明channel的格式:
	var 变量 chan 类型
*/

func main() {
	// var ch1 chan int   // 声明一个传递整型的通道
	// var ch2 chan bool  // 声明一个传递布尔型的通道
	// var ch3 chan []int // 声明一个传递int切片的通道

	// var ch chan int
	// fmt.Println(ch) // <nil>

	// 声明通道后需要使用make函数初始化之后才能够使用
	//   make(chan 元素类型, [缓冲大小])

	ch := make(chan int)
	// 将一个值发送到通道中
	ch <- 10 // 将10发送到ch中

	// 从通道中接收值
	// x := <-ch
	// <- ch  // 从ch中获取值, 忽略结果
	close(ch)


}

package main

import (
	"io"
	"log"
	"os"
	"strings"
)

/*
I/O操作也叫输入输出操作。其中I是指Input，O是指Output，用于读或者写数据的，有些语言中也叫流操作，是指数据通信的通道。

Golang 标准库对 IO 的抽象非常精巧，各个组件可以随意组合，可以作为接口设计的典范。

io包中提供I/O原始操作的一系列接口。

它主要包装了一些已有的实现，如 os 包中的那些，并将这些抽象成为实用性的功能和一些其他相关的接口。

由于这些接口和原始的操作以不同的实现包装了低级操作，客户不应假定它们对于并行执行是安全的。

**io库比较常用的接口有三个，分别是Reader，Writer和Closer。**
*/

/*
IO接口类型:

type Reader interface {
	Read(p []byte) (n int, err error)
}
*/

func main() {

	reader := strings.NewReader("zhangsan test123 123")
	// 每次读取4个字节
	p := make([]byte, 4)
	for {

		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				log.Printf("读完了:eof错误 :%d", n)
				break
			}
			log.Printf("其他错误:%v", err)
			os.Exit(2)
		}
		log.Printf("[读取到的字节数为:%d][内容:%v]", n, string(p[:n]))
	}

}

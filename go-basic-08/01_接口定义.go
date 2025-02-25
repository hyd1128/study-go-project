package main

import (
	"errors"
	"fmt"
)

/*
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
*/

// 定义writer接口
type Writer interface {
	Writer([]byte) error
}

// 定义一个文件写入器
type FileWriter struct {
	filename string
}

// 实现Write方法
func (f *FileWriter) Writer(data []byte) error {
	if len(data) == 0 {
		return errors.New("写入数据不能够为空")
	}

	fmt.Printf("将数据写入文件: %s \n", f.filename)
	return nil
}

func main() {
	var w Writer = &FileWriter{filename: "test.txt"}

	err := w.Writer([]byte("Hello, Go!"))
	if err != nil {
		fmt.Println("写入失败", err)
	} else {
		fmt.Println("写入成功")
	}

	err = w.Writer([]byte(""))
	if err != nil {
		fmt.Println("写入失败:", err)
	}
}

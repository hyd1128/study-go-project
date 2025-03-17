package main

import "fmt"

/*
接口的实现条件:
	1. 如果一个任意类型 T 的方法集为一个接口类型的方法集的超集，则我们说类型 T 实现了此接口类型。
	2. T可以是一个接口类型也可以是一个非接口类型
	3. 接口定义后,需要实现接口, 调用方才能够正确编译通过并使用接口

	接口需要遵循两条规则才能够让接口可用:
		1. 接口的方式与实现接口的类型方法格式一致
			在类型中添加与接口签名一致的方法就可以实现该方法。
			签名包括方法中的名称、参数列表、返回参数列表。
			也就是说，只要实现接口类型中的方法的名称、参数列表、返回参数列表中的任意一项与接口要实现的方法不一致，那么接口的这个方法就不会被实现。

			当类型无法实现接口的时候, 编译器会报错:
				1. 函数名不一致导致的报错
				2. 实现接口的方法签名不一致导致的报错
		2. 接口中所有的方法均被实现
*/

// 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
	/*
		该接口方法中, data interface{} 作为WriterData方法的参数, 表示该方法可以接收任意类型的数据
	*/

	// 空接口interface{} 的作用
	// 1. 接收任意类型的值, 由于go语言是静态类型的, 普通函数参数需要指定类型(如: string, int)
	// 2. 使用interface{} 可以让WriterData方法接收任何类型的data

	// 如果接口中新增加了一个方法, 则会报错
	// CanWrite() bool
}

// 定义文件结构, 用于实现DataWriter
type file struct {
}

// 实现DataWriter接口的WriterData方法
func (d *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriterData:", data)
	return nil
}

func main() {
	// 实例化file
	f := new(file)
	// 生命一个DataWriter的接口
	var writer DataWriter
	// 将接口赋值f, 也就是*file类型
	writer = f
	// 使用DataWriter接口进行数据写入
	err := writer.WriteData("data")
	if err != nil {
		return
	}

}

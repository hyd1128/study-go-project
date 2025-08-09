package main

import "fmt"

/*
go基本类型:
- bool
- string
- int（随系统，一般是占用4个字节）、int8(占一个字节)、int16(占两个字节)、int32(占4个字节)、int64（占8个字节）
- uint（无符号整数）、uint8、uint16、uint32、uint64、uintptr
- byte // `uint8 的别名`
- rune // `int32 的别名 代表一个 Unicode 码`
- float32、float64
- complex64、complex128

unicode和ascii的异同:
ASCII 只支持 128 个字符，而 Unicode 覆盖了全球所有语言的字符。
Unicode 的前 128 个字符与 ASCII 完全相同，确保了向后兼容。
UTF-8 编码方式使得 ASCII 字符在 Unicode 环境下仍然能用 1 字节表示，保持兼容性。
*/

// 当一个变量被声明后 系统会自动赋予它该值的零值
var age int

// 未指定类型 go编译器会自动推导类型
var level = 1

// 批量定义变量
var (
	a int
	b string
	c []float32
)

func main() {
	d := 100
	e := 200
	f, g := 100, 200

	fmt.Println(d, e, f, g)

	fmt.Println(age)
	fmt.Println(level)
	// 类型输出
	fmt.Printf("level变量的类型为: %T", level)
	// %d 整数占位符，%s 字符串占位符， %f 浮点数占位符(默认精度为6)
	fmt.Printf("%d,%s,%f", a, b, c)

	/*
		使用简短格式声明变量
		使用简短格式申明变量的限制:
			1. 定义变量，同时显式初始化
			2. 不能提供数据类型
			3. 只能用在函数内部
	*/
	i := 1
	fmt.Println(i)
	// 注意: 简短变量声明被广泛用于大部分的局部变量的声明和初始化，var 形式的声明语句往往用于需要显式指定变量类型的地方

	// go语言中的重复申明: 在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，即便其他变量名可能是重复声明的，编译器也不会报错

	a1, a2 := 2, 4
	fmt.Println(a1, a2)
	// 报错
	// a1, a2 := 2, 4

	// 正确写法
	a1, a3 := 1, 2
	fmt.Println(a1, a3)
}

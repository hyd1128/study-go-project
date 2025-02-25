package main

import "fmt"

/*
类型别名是 Go 1.9 版本添加的新功能，主要用于解决代码升级、迁移中存在的类型兼容性问题。
*/

// 将NewInt定义为int类型
type NewInt int

// 将int取一个别名叫IntAlias
type IntAlias = int

// 将a声明为NewInt类型
var a NewInt

func main() {
	/*
		类型别名:
			type TypeAlias = Type
		类型定义:
			定义Name为Type类型 ,定义之后 Name为一种新的类型
			type Name Type
	*/

	// 查看a的类型名 main.NewInt
	fmt.Printf("a type: %T\n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名 int
	// IntAlias 类型只会在代码中存在，编译完成时，不会有 IntAlias 类型。
	fmt.Printf("a2 type: %T\n", a2)

}

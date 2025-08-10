package main

import "fmt"

// func main() {
// // 常量再编译的时候就确定了值了，在运行的时候不会再分配内存地址
// const A = 1
// fmt.Printf("%v \n", A)
//
// var a int = 1
// var b string = "abc"
//
// // 取指针
// fmt.Printf("%p \n", &a)
// fmt.Printf("%p \n", &b)
//
// p1 := &a
// // 获取指针动态类型
// fmt.Printf("a的指针的类型: %T \n", &a)
// fmt.Printf("%v", *p1)

// 创建一个数组
// var array = [...]int{1, 2, 3, 4}
// fmt.Printf("%v, %T", array, array)

// 创建一个切片
// var slice []int
// fmt.Printf("%v, %T", slice, slice)

// 创建map的方法
// var mapOne = map[string]int{}
// var mapTwo map[string]int
// var mapThree = make(map[string]int, 10)
// var mapFour = make(map[string]int)

// var slice []string
// slice = append(slice, "a")
// slice = append(slice, "b")
// slice = append(slice, "c")
// slice = append(slice, "d")
//
// // 打印切片
// fmt.Printf("%v \n", slice)
//
// for i := 0; i < len(slice); i++ {
// 	mapOne[slice[i]] = i
// }
//
// fmt.Println(mapOne)
//
// // 遍历map
// for k, v := range mapOne {
// 	fmt.Println(k, v)
// }
// }

type Command struct {
	Name    string
	Var     *int
	Comment string
}

func generateCommand(name string,
	varRef *int,
	command string) *Command {
	// 字面量初始化
	return &Command{
		Name:    name,
		Var:     varRef,
		Comment: command,
	}
}

var version = 1

func main() {
	cmd := generateCommand("version",
		&version,
		"show version")

	println(cmd)
	fmt.Printf("值: %v, 类型: %T", cmd, cmd)
}

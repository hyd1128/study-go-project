package main

import (
	"fmt"
	"sync"
)

func main() {
	// // 创建一个int到int的映射
	// m := make(map[int]int)
	// // 开启一段并发代码
	// go func() {
	// 	// 不停地对map进行写入
	// 	for {
	// 		m[1] = 1
	// 	}
	// }()
	// // 开启一段并发代码
	// go func() {
	// 	// 不停地对map进行读取
	// 	for {
	// 		_ = m[1]
	// 	}
	// }()
	// // 无限循环, 让并发程序在后台执行
	// for {
	// }

	// sync.Map 不能使用 make 创建
	var scene sync.Map
	// 将键值对保存到sync.Map
	// sync.Map 将键和值以 interface{} 类型进行保存。
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	// 遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
	/*
		sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。
	*/

}

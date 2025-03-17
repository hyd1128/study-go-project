package main

import "fmt"

func main() {
	/*
		map 是一种无序的`键值对`的集合。
		map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
		map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，map 是无序的，我们无法决定它的返回顺序，这是因为 map 是使用 hash 表来实现的。
		map 是引用类型，可以使用如下方式声明
		[keytype] 和 valuetype 之间允许有空格。
		var mapname map[keytype]valuetype
	*/
	var mapList map[string]int
	var mapAssigned map[string]int

	mapList = map[string]int{"one": 1, "two": 2}
	mapAssigned = mapList
	// mapAssigned 是 mapList 的引用，对 mapAssigned 的修改也会影响到 mapList 的值。
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapList["one"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapList["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapList["ten"])

	/*
		map的另外一种创建方式
		make(map[keytype]valuetype)
		make(map[keytype]valuetype, cap)
		切记不要使用new创建map，否则会得到一个空引用的指针
		map 可以根据新增的 key-value 动态的伸缩，因此它不存在固定长度或者最大限制，但是也可以选择标明 map 的初始容量 capacity
	*/

}

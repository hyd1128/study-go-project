package main

import (
	"fmt"
)

func main() {
	scene := make(map[string]int)
	scene["cat"] = 66
	scene["dog"] = 88
	scene["bird"] = 99
	for k, v := range scene {
		fmt.Println(k, v)
	}
	// map是无序的，不要期望 map 在遍历时返回某种期望顺序的结果
	delete(scene, "dog")
	fmt.Println("----------------------")
	for k, v := range scene {
		fmt.Println(k, v)
	}
}

package main

import (
	"fmt"
)

/*
一个类型可以实现多个接口
多个类型也可以同时实现单个接口, 而接口间彼此独立, 不实现
*/

// Mover接口
type Mover interface {
	move()
}

// 狗可以动, 汽车也可以动
type pear struct {
	name string
}

type car struct {
	brand string
}

// dog类型实现Mover接口
func (d pear) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func main() {
	var x Mover
	var a = pear{name: "旺财"}
	var b = car{brand: "保时捷"}
	x = a
	x.move()
	x = b
	x.move()
}

package main

import "fmt"

/*
一个类型可以实现多个接口
一个类型可以同时实现多个接口, 而接口间彼此独立, 不实现
*/

// sayer接口
type Sayer interface {
	say()
}

// type Mover interface {
// 	move()
// }

// dog 既可以实现Sayer接口, 也可以实现Mover接口
type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()
}

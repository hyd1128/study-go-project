package main

import "fmt"

/*
变量交换，比如a=100，b=200，交换之后 a=200，b=100
*/

func main() {
	a := 10
	b := 20

	// method one
	// var c int
	// c = b
	// b = a
	// a = c
	// fmt.Println(a, b)

	// method two
	// a = a ^ b
	// b = b ^ a
	// a = a ^ b
	// fmt.Println(a, b)

	// method three
	b, a = a, b
	fmt.Printf("a=%d,b=%d", a, b)
}

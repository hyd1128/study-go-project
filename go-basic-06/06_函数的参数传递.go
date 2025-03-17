package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	var a, b int = 1, 2
	/*
		调用 swap() 函数
		&a 指向 a 指针, a变量的地址
		&b 指向 b 指针, b变量的地址
	*/

	swap(&a, &b)
	fmt.Println(a, b)
}

/*
多参数传递写法格式:
  func myfunc(args ...int) {    //0个或多个参数
  }

  func add(a int, args…int) int {    //1个或多个参数
  }

  func add(a int, b int, args…int) int {    //2个或多个参数
  }
*/

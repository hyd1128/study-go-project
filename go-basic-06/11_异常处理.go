package main

// func main() {
// 	testTwo()
// }
//
// func testTwo() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			println(err.(string)) // 将 interface{} 转型为具体类型。
// 		}
// 	}()
//
// 	panic("panic error!")
// }

// import "fmt"
//
// func testThree(x, y int) {
// 	var z int
//
// 	func() {
// 		defer func() {
// 			if recover() != nil {
// 				z = 0
// 			}
// 		}()
// 		panic("test panic")
// 		z = x / y
// 		return
// 	}()
//
// 	fmt.Printf("x / y = %d\n", z)
// }
//
// func main() {
// 	testThree(2, 1)
// }

// var ErrDivByZero = errors.New("division by zero")
//
// func div(x, y int) (int, error) {
// 	if y == 0 {
// 		return 0, ErrDivByZero
// 	}
// 	return x / y, nil
// }
//
// func main() {
// 	defer func() {
// 		fmt.Println(recover())
// 	}()
// 	switch z, err := div(10, 0); err {
// 	case nil:
// 		println(z)
// 	case ErrDivByZero:
// 		panic(err)
// 	}
// }

import "fmt"

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func main() {
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

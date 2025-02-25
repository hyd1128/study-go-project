package main

import (
	"errors"
	"fmt"
	"os"
)

// 未使用goto
// func main() {
// 	// err := firstCheckError()
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	exitProcess()
// 	// }
//
// 	err1 := secondCheckError()
// 	if err1 != nil {
// 		fmt.Println(err1)
// 		exitProcess()
// 	}
//
// 	fmt.Println("done")
// }
//
// func secondCheckError() interface{} {
// 	return errors.New("错误2")
// }
//
// func exitProcess() {
// 	// 退出
// 	os.Exit(1)
// }
//
// func firstCheckError() interface{} {
// 	return errors.New("错误1")
// }

// 使用goto
func main() {
	// err := firstCheckError()
	// if err != nil {
	// 	fmt.Println(err)
	// 	goto onExit
	// }
	err := secondCheckError()
	if err != nil {
		fmt.Println(err)
		goto onExit
	}
	fmt.Println("done")
	return
onExit:
	exitProcess()
}

func secondCheckError() interface{} {
	return errors.New("错误2")
}

func exitProcess() {
	fmt.Println("exit")
	// 退出
	os.Exit(1)
}

func firstCheckError() interface{} {
	return errors.New("错误1")
}

package main

import "fmt"

func main() {
	/*
		具体语法:
				switch var1 {
			    case val1:
			        ...
			    case val2:
			        ...
			    default:
			        ...
			}
	*/

	/* 定义局部变量 */
	var grade string = "B"
	var score int = 90

	switch score {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	// swtich后面如果没有条件表达式，则会对true进行匹配
	// swtich后面如果没有条件表达式，则会对true进行匹配
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	/*
		Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		那么如何做到执行完一个case之后，进入下一个case而不是跳出swtich呢？
	*/
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}

}

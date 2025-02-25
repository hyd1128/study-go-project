package main

import "fmt"

// Go语言支持两种浮点型数:
// 1. `float32` :范围 约1.4e-45 到 约3.4e38
// 2. `float64` :范围约4.9e-324 到 约1.8e308

func main() {
	// floatStr1 := 3.2
	// 保留小数点位数
	// fmt.Printf("%.2f\n", floatStr1)

	// 通常应该优先使用 float64 类型，因为 float32 类型的累计计算误差很容易扩散，并且 float32 能精确表示的正整数并不是很大
	// var f float32 = 1 << 24
	// fmt.Println(f == f+1) // true

	// 浮点数在声明的时候可以只写整数部分或者小数部分
	// var e = .71828 // 0.71828
	// var f = 1.     // 1
	// fmt.Printf("%.5f,%.1f", e, f)

	// 很小或很大的数最好用科学计数法书写, 通过 e 或 E 来指定指数部分
	var avogadro = 6.02214129e23 // 阿伏伽德罗常数
	var planck = 6.62606957e-34  // 普朗克常数
	fmt.Printf("%f,%.35f", avogadro, planck)
}

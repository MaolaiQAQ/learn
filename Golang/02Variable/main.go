package main

import "fmt"

func main() {
	// 1. 零值 (Zero Value) 演示
	var a int     // 默认 0
	var b float64 // 默认 0.0
	var c bool    // 默认 false
	var d string  // 默认 ""
	fmt.Printf("零值展示: [%v] [%v] [%v] [%q]\n", a, b, c, d)

	// 2. 类型转换 (Type Conversion)
	var i int = 42
	var f float64 = float64(i) // 必须显式转换：T(v)
	var u uint = uint(f)
	fmt.Printf("转换过程: int(%v) -> float64(%v) -> uint(%v)\n", i, f, u)

	// 3. 字符串拼接
	lang := "Go"
	phrase := lang + " 是一门艺术"
	fmt.Println(phrase)
}

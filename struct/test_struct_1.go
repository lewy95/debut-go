package main

import "fmt"

/*
类型定义与类型别名
*/
func main() {
	// 类型定义
	type MyInt int
	// i 为MyInt类型
	var i MyInt
	i = 100
	fmt.Printf("i: %v i: %T\n", i, i) // i: 100 i: main.MyInt

	// 类型别名定义
	type MyInt2 = int
	// i 其实还是int类型
	var j MyInt2
	j = 100
	fmt.Printf("j: %v j: %T\n", j, j) // j: 0 j: int
}

package main

import "fmt"

/*
测试匿名函数
*/
func main() {
	// 法一：
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	i := max(1, 2)
	fmt.Printf("i: %v\n", i) // i：2

	// 法二：自己执行
	func(a int, b int) {
		max := 0
		if a > b {
			max = a
		} else {
			max = b
		}
		fmt.Printf("max: %v\n", max) // max：2
	}(1, 2)
}

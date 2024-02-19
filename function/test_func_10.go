package main

import "fmt"

/*
递归测试：阶乘
*/
func multiple(n int) int {
	// 返回条件
	if n == 1 {
		return 1
	} else {
		// 自己调用自己
		return n * multiple(n-1)
	}
}

func main() {
	n := 5
	r := multiple(n)
	fmt.Printf("r: %v\n", r) // 120
}

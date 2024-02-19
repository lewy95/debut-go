package main

import "fmt"

func fibonacci(n int) int {
	// 退出点判断
	if n == 1 || n == 2 {
		return 1
	}
	// 递归表达式
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	r := fibonacci(10)
	fmt.Printf("r: %v\n", r) // 55
}

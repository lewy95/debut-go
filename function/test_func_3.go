package main

import "fmt"

/*
定义一个函数类型
*/
type fun func(int, int) int

func getSum(a int, b int) int {
	return a + b
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var f fun
	f = getSum
	s := f(1, 2)
	fmt.Printf("s: %v\n", s) // s: 3
	f = getMax
	m := f(3, 4)
	fmt.Printf("m: %v\n", m) // m: 4
}

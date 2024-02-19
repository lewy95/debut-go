package main

import "fmt"

/*
测试高阶函数，函数作为返回值
*/
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	add := cal("+")
	r := add(1, 2)
	fmt.Printf("r: %v\n", r)

	fmt.Println("-----------")

	sub := cal("-")
	r = sub(100, 50)
	fmt.Printf("r: %v\n", r)
}

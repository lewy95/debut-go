package main

import "fmt"

/*
闭包测试1
*/
// 声明了一个返回值为函数的函数
func addMethod() func(int) int {
	// 定义一个变量
	var x int
	// 定义一个函数
	return func(y int) int {
		x += y // 将变量和参数相加
		return x
	}
}

func addMethod2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	var f = addMethod()
	fmt.Println(f(10)) // 10
	fmt.Println(f(20)) // 10 + 20 = 30
	fmt.Println(f(30)) // 30 + 30 = 60
	fmt.Println("-----------")
	f1 := addMethod()   // 重新定义一个函数，作用域变化了
	fmt.Println(f1(40)) // 40
	fmt.Println(f1(50)) // 40 + 50 = 90

	fmt.Println("===========")

	var f3 = addMethod2(10)
	fmt.Println(f3(10)) // 10 + 10 = 20
	fmt.Println(f3(20)) // 20 + 20 = 40
	fmt.Println(f3(30)) // 40 + 30 = 70
	fmt.Println("----------")
	f4 := addMethod2(20)
	fmt.Println(f4(40)) // 20 + 40 = 60
	fmt.Println(f4(50)) // 60 + 50 = 110
}

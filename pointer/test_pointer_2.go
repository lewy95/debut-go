package main

import "fmt"

/**
指向数组的指针
*/
func main() {
	// 声明一个数组
	a := []int{1, 3, 5}
	// 声明一个指针数组
	var ptr [3]*int
	fmt.Println(ptr) // 可见指针数组的初始值为[<nil> <nil> <nil>]
	for i := 0; i < 3; i++ {
		// 整数地址赋值给指针数组
		ptr[i] = &a[i]
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i]) // *ptr[i]就是打印出相关指针的值了
	}
	// 执行结果：
	// a[0] = 1
	// a[1] = 3
	// a[2] = 5
}

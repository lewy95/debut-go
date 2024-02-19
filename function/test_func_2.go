package main

import "fmt"

func testDelivery(a int) {
	a = 200
	fmt.Printf("a1: %v\n", a)
}

func testDelivery2(b []int) {
	// 修改了切片中某个值
	b[0] = 100
}

func main() {
	// 情况一：
	a := 100
	testDelivery(a)
	fmt.Printf("a: %v\n", a)
	// 运行结果：
	// a1: 200
	// a: 100
	// 可见，调用函数testValueDelivery后，a的值并没有被改变，说明参数传递是拷贝了一个副本，也就是拷贝了一份新的内容进行运算。

	// 情况二：
	b := []int{1, 2}
	testDelivery2(b)
	fmt.Printf("b: %v\n", b)
	// 运行结果：
	// b: [100 2]
	// 对于 map、slice、interface、channel这些数据类型本身就是指针类型的，所以就算是拷贝传值也是拷贝的指针，
	// 拷贝后的参数仍然指向底层数据结构，所以修改它们可能会影响外部数据结构的值。
}

package main

import "fmt"

/*
测试高阶函数，函数作为参数
*/
func sayHello(name string) {
	fmt.Printf("Hello,%s", name)
}

func testHigh(name string, f func(string)) {
	f(name)
}

func main() {
	testHigh("fish", sayHello)
}

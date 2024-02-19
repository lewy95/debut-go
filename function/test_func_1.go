package main

import "fmt"

// 单返回值定义：法一
func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

// 单返回值定义：法二
func sum2(a int, b int) int {
	ret := a + b
	return ret
}

// 单返回值定义：法三
func sum3(a int, b int) int {
	return a + b
}

func getPerson() (name string, age int) {
	name = "fish"
	age = 26
	return name, age
}

func getPerson2() (name string, age int) {
	name = "fish"
	age = 26
	return // 等价于return name, age
}

// 不使用定义的返回值名称，这种做法不建议使用
func getPerson3() (name string, age int) {
	n := "fish"
	a := 26
	return n, a
}

func testArgs(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}
func testArgs2(name string, age int, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	sum := sum(1, 2)
	fmt.Printf("sum: %v\n", sum)

	// name, age := getPerson()
	// _, age := getPerson2() // 不用必须用 _ 替换
	name, age := getPerson3()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

	testArgs(1, 2)
	testArgs2("fish", 26, 1, 2, 3)

}

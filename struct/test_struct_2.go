package main

import "fmt"

type Person struct {
	id    int
	name  string
	age   int
	email string
}

type Person2 struct {
	id, age     int
	name, email string // 形同类型的可以合并到一行
}

func main() {
	var fish Person
	fmt.Printf("fish: %v\n", fish) // fish: {0  0 } 没有赋值前都是默认初始值
	lewy := Person{}
	fmt.Printf("lewy: %v\n", lewy)

	// 赋值：法一：
	lewy.id = 1
	lewy.name = "tom"
	lewy.age = 20
	lewy.email = "tom@gmail.com"
	fmt.Printf("lewy: %v\n", lewy) // lewy: {1 tom 20 tom@gmail.com}

	// 赋值：法二：
	mm := Person{1, "mm", 20, "mm@gmail.com"}
	fmt.Printf("mm: %v\n", mm)

	// 赋值：法三：使用键值对进行初始化
	yy := Person{
		id:    1,
		name:  "yy",
		age:   20,
		email: "yy@gmail.com",
	}
	fmt.Printf("yuan: %v\n", yy)
}

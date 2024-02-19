package main

import "fmt"

/**
结构体指针
*/
func main() {

	// 属性指针
	var name string
	name = "tom"

	var pName *string                 // p_name 指针类型
	pName = &name                     // &name 取name地址
	fmt.Printf("name: %v\n", name)    // 输出指针地址
	fmt.Printf("p_name: %v\n", pName) // 输出指针指向的内容值
	fmt.Printf("*p_name: %v\n", *pName)

	// 结构体指针
	type Person struct {
		id   int
		name string
	}

	var tom = Person{1, "tom"}

	var pPerson *Person
	pPerson = &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %p\n", pPerson)
	fmt.Printf("*p_person: %v\n", *pPerson)
}

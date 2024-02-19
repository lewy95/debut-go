package main

import "fmt"

func main() {
	// 匿名结构体
	var person struct {
		id   int
		name string
	}
	person.id = 1
	person.name = "fish"
	fmt.Printf("person: %v\n", person) // person: {1 fish}
}

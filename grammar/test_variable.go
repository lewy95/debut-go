package grammar

import "fmt"

func TestVariable() {
	// var name string = "fish"
	var name = "fish" // go 支持类型推导
	var age = 26
	var sex, addr = "female", "heart" // 支持批量初始化
	// 短变量声明方式，只能在函数内声明
	phone := "18611111111"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(sex)
	fmt.Println(addr)
	fmt.Println(phone)

	newName, _ := getNameAndAge()
	fmt.Printf("name: %v\n", newName)
}

func getNameAndAge() (string, int) {
	return "fish", 26
}

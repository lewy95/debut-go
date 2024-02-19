package grammar

import "fmt"

func TestIf() {

	// 声明变量和条件
	var age = 20
	if age > 18 {
		fmt.Println("你是成年人")
	}

	// 初始变量可以声明在布尔表达式里面，注意它的作用域
	if age := 26; age > 18 {
		fmt.Println("你是成年人")
	}

}

func TestIfElseIf() {
	score := 80
	if score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score > 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}

}

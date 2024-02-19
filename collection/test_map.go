package collection

import "fmt"

func TestMap() {
	// 声明方式一：
	m1 := make(map[string]string)
	m1["name"] = "fish"
	m1["age"] = "26"
	m1["email"] = "fish@gmail.com"

	fmt.Printf("m1: %v\n", m1)

	// 声明方式二：
	m2 := map[string]string{
		"name":  "lewy",
		"age":   "26",
		"email": "lewy@gmail.com",
	}

	fmt.Printf("m2: %v\n", m2)

	// 判断 key 是否存在
	// v, ok := m1["address"]
	v, ok := m1["name"]
	if ok {
		fmt.Println("键存在")
		fmt.Println(v)
	} else {
		fmt.Println("键不存在")
	}

	// 遍历 map
	// 法一：
	for key := range m2 {
		fmt.Println(key)
	}

	// 法二：
	for key, value := range m2 {
		fmt.Println(key + ":" + value)
	}
}

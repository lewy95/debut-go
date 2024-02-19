package grammar

import "fmt"

func testFor() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	// 支持将初始值写到外层
	j := 1
	for ; j <= 10; j++ {
		fmt.Printf("i: %v\n", j)
	}
}

// Go 语言永久循环，相当于 while
func testForever() {
	for {
		fmt.Println("我一直在执行~")
	}
}

func testForRange() {
	// 定一个字符串
	var s = "多课网，go教程"
	for i, v := range s {
		fmt.Printf("i: %d, v: %c\n", i, v)
	}

	// 定义一个数组
	var a = [5]int{1, 2, 3, 4, 5}
	// 可以打印索引和值
	for i, v := range a {
		fmt.Printf("i: %d, v: %v\n", i, v)
	}

	// 定一个map
	m := make(map[string]string)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "tom@gmail.com"
	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
}

package grammar

import "fmt"

func testBreak() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 退出循环
		}
		fmt.Printf("i: %v\n", i)
	}
	// 结果：0 1 2 3 4

	// 跳出到指定循环
MyLabel:
	for i := 0; i < 10; i++ {
		if i == 5 {
			break MyLabel
		}
		fmt.Printf("%v\n", i)
	}
	// 结果：0 1 2 3 4
}

func testContinue() {

	// MY_LABEL:
	for i := 0; i < 5; i++ {
	MyLabel:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue MyLabel
			}
			fmt.Printf("i=%d,j=%d\n", i, j)
		}
	}

}

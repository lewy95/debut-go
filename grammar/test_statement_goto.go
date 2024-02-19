package grammar

import "fmt"

func TestGoto(age int) {
	if age >= 18 {
		goto LABEL1
	} else {
		// goto LABEL2 // // goto 后面的语句无法执行
		fmt.Printf("bad")
	}

LABEL1:
	fmt.Printf("good")

	// 不管是不是条件中跳转过来的，标签中的代码也会因为上到下而执行
	//LABEL2:
	//	for i := 0; i < 10; i++ {
	//		fmt.Printf("%v\n", i)
	//	}
}

// TestGoto2 跳出双重循环
func TestGoto2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				// 使用 break，只会跳出一层循环，无法跳出两层循环，但goto可以
				goto LABEL1
			}
		}
	}
LABEL1:
	fmt.Println("label1")
}

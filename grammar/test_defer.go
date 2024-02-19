package grammar

import "fmt"

func TestDefer() {
	fmt.Println("start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("end")
}

// 执行结果：
// start
// end
// step3
// step2
// step1

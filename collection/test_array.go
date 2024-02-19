package collection

import "fmt"

func TestArray() {
	// 定义空数组
	var intArr [3]int    // 定义一个int类型的数组a，长度是3
	var strArr [2]string // 定义一个字符串类型的数组s，长度是2
	var boolArr [2]bool

	fmt.Printf("intArr: %v\n", intArr)   // intArr: [0 0 0]
	fmt.Printf("strArr: %v\n", strArr)   // strArr: ["" ""]
	fmt.Printf("boolArr: %v\n", boolArr) // boolArr: [false false]

	// 定义时赋值
	var intArr2 = [3]int{1, 2, 3}
	var strArr2 = [2]string{"lewy", "fish"}
	var boolArr2 = [2]bool{true, false}
	intArrInfer := [2]int{1, 2} // 支持类型推断

	fmt.Printf("intArr2: %v\n", intArr2)
	fmt.Printf("strArr2: %v\n", strArr2)
	fmt.Printf("boolArr2: %v\n", boolArr2)
	fmt.Printf("intArrInfer: %v\n", intArrInfer)

	// 省略长度定义
	var intArr3 = [...]int{1, 2, 3}
	var strArr3 = [...]string{"lewy", "fish"}
	var boolArr3 = [...]bool{true, false}
	intArrInfer2 := [...]int{1, 2} // 类型推断

	fmt.Printf("intArr3: %v\n", intArr3)
	fmt.Printf("strArr3: %v\n", strArr3)
	fmt.Printf("boolArr3: %v\n", boolArr3)
	fmt.Printf("intArrInfer2: %v\n", intArrInfer2)

	// 通过指定索引值来初始化，长度取决于最大的索引，没指定的项都为初始值
	var intArr4 = [...]int{0: 1, 2: 2}
	var strArr4 = [...]string{1: "lewy", 2: "muller"}
	var boolArr4 = [...]bool{2: true, 5: false}
	intArrInfer3 := [...]int{1, 2} // 类型推断

	fmt.Printf("intArr4: %v\n", intArr4)   // [1 0 2]
	fmt.Printf("strArr4: %v\n", strArr4)   // [ lewy fish] 0索引处有一个空串
	fmt.Printf("boolArr4: %v\n", boolArr4) // [false false true false false false]
	fmt.Printf("intArrInfer3: %v\n", intArrInfer3)
}

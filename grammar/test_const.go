package grammar

import (
	"fmt"
	"unsafe"
)

func TestConst() {
	// 常量声明了，可以不被使用
	const PI float32 = 3.14
	// const PI2 = 3.1415 // 类型也可以省略
	// 批量声明，通用用于枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值，但这些函数必须是内置函数，否则编译不过
	const (
		str    = "abc"
		length = len(str)
		size   = unsafe.Sizeof(str)
	)

	fmt.Println(PI)
	fmt.Println(length)
	fmt.Println(size)

	//testIota()
	//testIota2()
	testIota3()
}

func testIota() {
	const (
		a1 = iota // 0
		a2 = iota // 1
		a3 = iota // 2
	)
	fmt.Println(a1) // 0
	fmt.Println(a2) // 1
	fmt.Println(a3) // 2
}

/** 跳过值 */
func testIota2() {
	const (
		b1 = iota // 0
		_         // 1 _表示跳过
		b2 = iota // 2
	)
	fmt.Println(b1) // 0
	fmt.Println(b2) // 2
}

/** 插入值 */
func testIota3() {
	const (
		c1 = iota // 0
		c2 = 99   // 99 中间插入一个值
		c3 = iota // 2
	)
	fmt.Println(c1) // 0
	fmt.Println(c2) // 99
	fmt.Println(c3) // 2
}

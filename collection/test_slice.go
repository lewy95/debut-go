package collection

import "fmt"

func TestSlice() {
	// 定义空切片
	var names []string
	var numbers []int
	fmt.Printf("names: %v\n", names)
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Println(names == nil)   // true
	fmt.Println(numbers == nil) // true

	// 定义时赋值
	var names2 = []string{"tom", "kite"}
	var numbers2 = []int{1, 2, 3}
	// 通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量
	fmt.Printf("len: %d cap: %d\n", len(names2), cap(names2))
	fmt.Printf("len: %d cap: %d\n", len(numbers2), cap(numbers2))

	var strSlice = make([]string, 2, 3)
	fmt.Printf("len: %d cap: %d\n", len(strSlice), cap(strSlice))

	// 使用数组初始化
	intArr := [...]int{1, 2, 3}
	intSlice := intArr[:]
	fmt.Printf("intSlice: %v\n", intSlice) // [1,2,3]

	// 使用数据部分元素初始化
	arr := [...]int{1, 2, 3, 4, 5, 6}
	s1 := arr[2:5]
	fmt.Printf("s1: %v\n", s1)
	s2 := arr[2:]
	fmt.Printf("s2: %v\n", s2)
	s3 := arr[:3]
	fmt.Printf("s3: %v\n", s3)

	// 切片遍历，法一：
	ts1 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(ts1); i++ {
		fmt.Printf("ts1[%d]: %v\n", i, ts1[i])
	}

	// 切片遍历，法二：
	ts2 := []int{1, 2, 3, 4, 5}
	for i, v := range ts2 {
		fmt.Printf("ts2[%d]: %v\n", i, v)
	}
}

// TestSliceAppend 切片添加元素
func TestSliceAppend() {
	var s1 []int
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3, 4, 5)   // 添加多个元素
	fmt.Printf("s1: %v\n", s1) // s1: [1 2 3 4 5]

	s3 := []int{3, 4, 5}
	s4 := []int{1, 2}
	s4 = append(s4, s3...)     // 添加另外一个切片
	fmt.Printf("s4: %v\n", s4) //s4: [1 2 3 4 5]

	s4[3] = 118
	fmt.Printf("s4: %v\n", s4) //s4: [1 2 3 4 5]

}

// TestSliceRemove 切片删除元素
func TestSliceRemove() {
	s1 := []int{1, 2, 3, 4, 5}
	// 删除索引为2的元素
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1) // s1: [1 2 4 5]
}

// TestSliceCopy 拷贝切片
func TestSliceCopy() {
	s1 := []int{1, 2, 3}
	s2 := s1
	s1[0] = 100                // 会连带修改
	fmt.Printf("s1: %v\n", s1) // s1: [100 2 3]
	fmt.Printf("s2: %v\n", s2) // s2: [100 2 3]
	fmt.Println("----------")

	s3 := make([]int, 3)
	copy(s3, s1) // 使用copy，值不会连带修改
	s1[0] = 1
	fmt.Printf("s1: %v\n", s1) // s1: [1 2 3]
	fmt.Printf("s3: %v\n", s3) // s3: [100 2 3]
}

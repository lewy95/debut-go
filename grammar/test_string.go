package grammar

import (
	"bytes"
	"fmt"
	"strings"
)

func TestString() {
	var str1 = "hello world"
	var html = `
		<html>
		<head><title>hello golang</title>
		</html>
			`

	fmt.Printf("str1: %v\n", str1)
	fmt.Printf("html: %v\n", html)
}

// TestConnectString 字符串连接
func TestConnectString() {
	// 法一：使用 + 号拼接
	name := "fish"
	age := "20"
	msg := name + " " + age
	fmt.Printf("msg: %v\n", msg)
	fmt.Println("-------------")

	// 法二：fmt.Sprintf 打印输出
	msg2 := fmt.Sprintf("%s##%s", name, age)
	fmt.Printf("msg: %v\n", msg2)
	fmt.Println("-------------")

	// 法三：使用strings模块下的Join方法
	msg3 := strings.Join([]string{name, age}, ",")
	fmt.Printf("msg: %v\n", msg3)
	fmt.Println("-------------")

	// 法四：使用bytes.Buffer库，这个比较理想，可以当成可变字符使用，对内存的增长也有优化，如果能预估字符串的长度，还可以用 buffer.Grow() 接口来设置 capacity
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String())
	fmt.Println("-------------")

}

// TestGetString 字符串切片方式
func TestGetString() {
	str := "hello world"
	n := 3
	m := 5
	fmt.Println(str[n])   // 获取字符串索引位置为n的原始字节
	fmt.Println(str[n:m]) // 截取得字符串索引位置为 n 到 m-1 的字符串
	fmt.Println(str[n:])  // 截取得字符串索引位置为 n 到 len(s)-1 的字符串
	fmt.Println(str[:m])  // 截取得字符串索引位置为 0 到 m-1 的字符串
}

// TestStringUseful 字符串常用方法
func TestStringUseful() {
	s := "hello world！"
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("strings.Split(s, \"\"): %v\n", strings.Split(s, " "))
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "hello"))
	fmt.Printf("strings.HasSuffix(s, \"world！\"): %v\n", strings.HasSuffix(s, "world！"))
	fmt.Printf("strings.Index(s, \"l\"): %v\n", strings.Index(s, "l"))
	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))
}

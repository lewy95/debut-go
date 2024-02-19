package main

import (
	"fmt"
	"strings"
)

/*
闭包测试2
*/
// 定义一个返回值为函数的函数
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test"))
}

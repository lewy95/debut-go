package main

import "fmt"

/**
指针基础
*/
func main() {
	var value = 20                           /* 声明实际变量 */
	var addr *int                            /* 声明指针变量 */
	addr = &value                            /* 指针变量的存储地址 */
	fmt.Printf("value 变量的地址是: %x\n", &value) // value 变量的地址是: c00012a008
	/* 指针变量的存储地址 */
	fmt.Printf("addr 变量储存的指针地址: %x\n", addr) // addr 变量储存的指针地址: c00012a008
	/* 使用指针访问值 */
	fmt.Printf("*addr 变量的值: %d\n", *addr) // *addr 变量的值: 20
}

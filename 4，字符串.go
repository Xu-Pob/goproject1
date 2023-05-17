package main

import (
	"fmt"
	"unsafe"
)

// 1,所有的字符串都是stringstruct结构体 ，两个指针长度
// a := make(chan struct{}) 传递消息
// len(s)获取的是字节长度
// 中文3英文1 utf-8
// 用for循环遍历字符、
// for _, char := range s
func main() {
	//1,所有的字符串都是stringstruct结构体 ，两个指针长度
	println(unsafe.Sizeof("张三"))
	println(unsafe.Sizeof("张三风"))
	m := map[string]struct{}{}
	m["a"] = struct{}{}
	//a := make(chan struct{}) 传递消息
	s := "我是Pobo"
	println(len(s)) //3+3+4    中文3英文1 utf-8
	//在这里打印的是字节
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	fmt.Println("______________")
	//用for循环遍历字符、
	for _, char := range s {
		fmt.Printf("char:%d ,%c\n", char, char)
	}
}

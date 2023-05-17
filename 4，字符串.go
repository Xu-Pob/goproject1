package main

import "unsafe"

func main() {
	//1,所有的字符串都是stringstruct结构体 ，两个指针长度
	println(unsafe.Sizeof("张三"))
	println(unsafe.Sizeof("张三风"))
}

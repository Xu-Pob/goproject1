package main

import (
	"fmt"
	"unsafe"
)

type K struct {
}
type F struct {
	num1 K
	num2 int
}

func main() {
	a := K{}
	fmt.Println(unsafe.Sizeof(int(0))) //跟随系统字长
	fmt.Println(unsafe.Sizeof(a))      //空结构体是0
	fmt.Printf("%p\n", &a)             //同一个空结构体都是同一个地址
	c := F{}
	fmt.Printf("%p\n", c.num1)
	fmt.Printf("%p", &c.num1)
}

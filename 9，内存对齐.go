package main

import (
	"fmt"
	"unsafe"
)

// S1 内存对齐   对齐系数 unsafe.Alignof()  变量的内存长度必须被对齐系数整除
// 结构体内部对齐，每个成员的偏移量是自身大小与其对齐系数较小值的倍数
// 结构体长度填充：对齐系统字长
// 结构体长度是最大成员长度与系统字长较小的整数倍
// 结构体的对齐系数是成员的最大对齐系数  ，    可以通过调整成员顺序，节约空间
func main() {
	s1 := S1{}
	s2 := S2{}
	println(unsafe.Sizeof(s1))
	println(unsafe.Alignof(s1))

	println(unsafe.Sizeof(s2))
	println(unsafe.Alignof(s2))

	fmt.Printf("bool size:%d align:%d\n", unsafe.Sizeof(bool(true)), unsafe.Alignof(bool(true)))
	fmt.Printf("byte size:%d align:%d\n", unsafe.Sizeof(byte(0)), unsafe.Alignof(byte(0)))
	fmt.Printf("int8 size:%d align:%d\n", unsafe.Sizeof(int8(0)), unsafe.Alignof(int8(0)))
	fmt.Printf("int16 size:%d align:%d\n", unsafe.Sizeof(int16(0)), unsafe.Alignof(int16(0)))
	fmt.Printf("int32 size:%d align:%d\n", unsafe.Sizeof(int32(0)), unsafe.Alignof(int32(0)))
	fmt.Printf("string size:%d align:%d\n", unsafe.Sizeof("1111"), unsafe.Alignof("1111"))
	fmt.Printf("[]int32 size:%d align:%d\n", unsafe.Sizeof([]int32{}), unsafe.Alignof([]int32{}))
}

type S1 struct {
	num1 int32
	num2 int32
}
type S2 struct {
	num1 int16
	num2 int32
}

// User 给出size 和 align大小 ？    4
type User struct {
	A int32
	B []int32
	C string
	D bool
	e struct{}
}

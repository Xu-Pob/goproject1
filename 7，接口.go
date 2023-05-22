package main

import (
	"fmt"
)

// 接口值的底层表示：runtime.iface          unsafe.Pointer万能指针
// 记录了数据的地址
// 也记录了接口类型信息和实现的方法=》做类型断言
// 类型断言：是一个使用在接口值上的操作：a,将接口值转为其他类型值（实现或者兼容接口）b,可以配合switch做类型判断
func main() {
	var c Car = Truct{}
	t := c.(Truct)
	//tt := c.(TrafficTool)
	fmt.Println(t)
	//fmt.Println(tt)
	switch c.(type) {
	//case Car:
	//	fmt.Println("car")
	//case Truct:
	//	fmt.Println("Truct")
	case TrafficTool:
		fmt.Println("TrafficTool")
	}
}

type Car interface {
	drive()
}
type TrafficTool interface {
	drive()
	fly()
}
type Truct struct {
	Model string
}

func (c Truct) drive() {
	fmt.Println(c.Model)
}

//func (c Truct) fly() {
//	fmt.Println(c.Model)
//}

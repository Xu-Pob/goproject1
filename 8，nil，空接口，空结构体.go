package main

import "fmt"

// 1，nil=>空  ，不一定是空指针   builtin.go 下面有对应描述   var nil Type
//Type must be a pointer, channel, func, interface, map, or slice type, 这六种类型的零值
//类型不同无法比较
//2，空结构体  ，值不是nil,但是都相同（zerobase）
//3,空接口    a,空接口零值是nil，一旦有了类型信息就不是nil

func main() {
	var a *int
	fmt.Println(a == nil)
	var b map[int]int
	fmt.Println(b == nil)

	//fmt.Println(a==b)   //类型不同无法比较
	c := KK{}
	d := LL{}
	e := LL{}
	fmt.Printf("%p\n", &c)
	fmt.Printf("%p\n", &d)
	fmt.Printf("%p\n", &e)

	//空接口
	var f interface{}
	var g interface{}
	println(f == nil)
	println(g == nil)
	println(f == g)
	//
	var h interface{}
	var i *int
	h = i
	fmt.Println(i == nil) //  h:eface{}    类型有值， data没 ， 还是有值，不为nil
	fmt.Println(h == nil)
}

type KK struct {
}
type LL struct {
}

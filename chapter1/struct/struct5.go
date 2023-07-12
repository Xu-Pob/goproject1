package main

import (
	"fmt"
)

type People struct {
	Name string
}

func (p People) SetName(name string) {
	p.Name = name
}

func (p People) Hello() {
	fmt.Println("Hi," + p.Name + ", I am Barry")
}

// 定义组合的结构体
type Student1 struct {
	People
}

// 定义组合的结构体
type Student2 struct {
	*People
}

func main() {
	name := "Barry"
	//定义
	a := People{"Jar"}
	//方法是指针引用，传递set赋值成功，否则是副本赋值
	a.SetName(name)
	a.Hello()

	//结构体组合
	b := Student1{People{"A"}}
	b.People.Hello()
	b.SetName("a")
	b.Hello()

	c := Student2{&People{"A"}}
	c.People.Hello()
	c.SetName("a")
	//?
	c.Hello()
}

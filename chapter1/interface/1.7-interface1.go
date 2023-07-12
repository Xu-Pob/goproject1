package main

import "fmt"

type Programmer1 struct {
	Name string
}
type SkillInterface interface {
	Write()
}

func (p Programmer1) Write() {
	fmt.Println("program Write()")
}
func main() {
	var pro Programmer1 //结构体变量实现了 Write()方法，实现了 SkillInterface接口
	var a SkillInterface = pro
	a.Write()
}

package main

import (
	"fmt"
)

func main() {
	//fmt.Println()
	var ap *int
	a := 12
	ap = &a
	fmt.Println(ap)
	fmt.Println(*ap)

	println("hello,go!")
	p := persion{name: "pobo", age: 40, gender: "man"}

	fmt.Println(p)
	p.describle()
	p.setAge(18)
	p.describle()
	dd1 := dog{Wake: true, Type: "typedog"}
	//接口
	var dd animal
	dd = dd1
	fmt.Println(dd1.Type)
	dd = cat{Sound: "Sssss", Type: "typecat"}
	fmt.Println(dd.decription())

}
func (p *persion) describle() {
	fmt.Printf("%v is %v year old\n", p.name, p.age)
}
func (p *persion) setAge(age int) {
	p.age = age
}
func (p persion) setGender(gender string) {
	p.gender = gender
}

type persion struct {
	name   string
	age    int
	gender string
}
type animal interface {
	decription() string
}
type dog struct {
	Type string
	Wake bool
}
type cat struct {
	Type  string
	Sound string
}

func (d dog) decription() string {
	return fmt.Sprintf("Wake:%v", d.Wake)
}
func (c cat) decription() string {
	return fmt.Sprintf("Sound:%v", c.Sound)
}

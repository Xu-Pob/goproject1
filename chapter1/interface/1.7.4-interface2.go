package main

import "fmt"

// Message 是一个定义了通知类行为的接口
type Message interface {
	sending()
}

// 定义Programmer 及Programmer.notify方法
type Programmer struct {
	name  string
	phone string
}

func (u *Programmer) sending() {
	fmt.Printf("Sending Programmer phone to %s<%s>\n", u.name, u.phone)
}

// 定义Student及Student.message方法
type Student struct {
	name  string
	phone string
}

func (a Student) sending() {
	fmt.Printf("Sending Student phone to %s<%s>\n", a.name, a.phone)
}

func main() {
	bile := Programmer{"Jack", "123"}
	lisa := Student{"Wade", "wasd123"}
	sendMessage(&bile)
	sendMessage(lisa)
}

// 接口作为参数
func sendMessage(message Message) {
	message.sending()
}

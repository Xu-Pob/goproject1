package main

import (
	"fmt"
	"time"
)

func Hello(c chan string) {
	//name := <-c //从通道获取数据
	if val, ok := <-c; ok {
		fmt.Println(val)
	}
	//fmt.Println(name)
}
func main() {
	//声明一个字符串类型的变量
	ch := make(chan string)
	//开启一个 goroutine
	go Hello(ch)

	//发送数据到通道
	ch <- `Hello`

	//关闭通道
	close(ch)
	//TestChan()
}

func TestChan() {
	ch := make(chan int)

	go func() {
		defer close(ch) // 在匿名函数中使用defer语句关闭通道

		// 在这里可以进行一些操作，发送数据到通道或者从通道接收数据
		time.Sleep(5 * time.Second) // 假设需要执行一段时间的操作

		// 示例操作：向通道发送一些数据
		ch <- 10
		ch <- 20
	}()

	// 等待接收数据
	for {
		if val, ok := <-ch; ok {
			fmt.Println(val)
		} else {
			break
		}
	}

}

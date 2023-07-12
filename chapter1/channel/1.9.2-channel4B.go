package main

import (
	"fmt"
	"time"
)

// 单向发送通道 zuowei参数
func sendMessage(ch chan<- string) {
	ch <- "Hello"
	ch <- "world"
	close(ch)
}

// 单向只读通道 zuowei参数
func receiveMessage(ch <-chan string) {
	for msg := range ch {
		time.Sleep(time.Second)
		fmt.Println("Recived", msg)
	}

}
func main() {
	var c = make(chan string)
	go sendMessage(c)

	receiveMessage(c)
}

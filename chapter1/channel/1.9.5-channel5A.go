package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second) // 创建一个每秒触发一次的定时器

	go func() {
		for range ticker.C {
			fmt.Println("Tick!")
		}
	}()

	time.Sleep(5 * time.Second) // 睡眠 5 秒钟，等待定时器执行

	ticker.Stop() // 停止定时器
	fmt.Println("Ticker stopped.")
}

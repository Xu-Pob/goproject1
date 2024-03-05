package main

import (
	"fmt"
	"time"
)

// 在这种情况下，可能是因为goroutines的执行速度非常快，
// 每个goroutine都能够在其他goroutine对b进行加1操作之前完成自己的加1操作和打印操作，
// 导致输出结果是固定的1到10。这种情况下是一种幸运的情况，实际上在并发编程中应该避免依赖于这种偶然性
func main() {
	b := 0
	for i := 0; i < 10; i++ {
		go func(index int) {
			b += 1
			fmt.Println(b)
		}(i)
	}
	time.Sleep(time.Second)
}

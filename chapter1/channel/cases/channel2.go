package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest(r chan<- int32) {
	time.Sleep(time.Second * 3) // 模拟一个工作负载
	r <- rand.Int31n(100)

}

func sumSquares(a, b int32) int32 {

	return a*a + b*b
}

// 带有缓冲的 channel 经常用于处理批量的数据、控制并发访问资源、实现生产者-消费者模型等场景
func main() {
	rand.Seed(time.Now().UnixNano())
	ra, rb := make(chan int32), make(chan int32)
	go longTimeRequest(ra)
	go longTimeRequest(rb)

	fmt.Println(sumSquares(<-ra, <-rb))

	results := make(chan int32, 2) // 缓冲与否不重要
	go longTimeRequest(results)
	go longTimeRequest(results)
	fmt.Println(sumSquares(<-results, <-results))
}

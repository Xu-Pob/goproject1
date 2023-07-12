package main

import (
	"fmt"
	"time"
)

// a,time定时器; b,select监听通道；
//select 语句的使用场景包括同时监听多个通道的读取或写入操作，
//以及多个 goroutine 之间的通信等。通过 select 语句，你可以实现非阻塞式的并发操作，
//编写更灵活和高效的代码。

type Request interface{}

const (
	RateLimitPeriod = time.Minute
	RateLimit       = 100
)

func handle(request Request) {
	fmt.Println(request.(int))
}
func handleRequest(requests <-chan Request) {
	//协程里面 定时qi
	quotas := make(chan time.Time, RateLimit)
	//定一个匿名协程， 每分钟最多执行次数发送消息，阻塞
	go func() {
		timeTick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer timeTick.Stop()
		for t := range timeTick.C {

			select {
			case quotas <- t:
				//default:
			}
		}
	}()
	for request := range requests {
		//释放阻塞 执行打印
		<-quotas

		handle(request)
	}
}
func main() {
	requests := make(chan Request)
	go handleRequest(requests)
	defer close(requests)
	for i := 0; i < 10; i++ {
		requests <- i
	}
	fmt.Println("结束了")
}

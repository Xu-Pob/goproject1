package main

import (
	"fmt"
	"sync"
	"time"
)

// 我们加锁,避免竞态访问，确保多协程都执行完毕
// 在Go语言中，sync.WaitGroup结构体中有一个内部计数器，用于跟踪正在等待的goroutine的数量。当调用WaitGroup的Add方法来增加计数器时，计数器会增加；当调用WaitGroup的Done方法来减少计数器时，计数器会减少；当调用Wait方法时，主goroutine会被阻塞，直到计数器减少到0。
// 如果计数器的值小于0，会导致sync: negative WaitGroup counter的错误。这通常是因为Done方法的调用次数多于Add方法的调用次数，或者在Wait方法之前计数器已经小于0。
func main() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}
	b := 0
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(index int) {
			mutex.Lock()
			b += 1
			fmt.Println(b)
			defer func() {
				mutex.Unlock()
				wait.Done()
			}()
		}(i)
	}
	time.Sleep(time.Second)
	wait.Wait()
}

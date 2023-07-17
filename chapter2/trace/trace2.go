package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

// 使用命令： go tool trace -http=8080 myTrace.out   跟踪这个任务相关的信息

func main() {
	runtime.GOMAXPROCS(1)
	f, _ := os.Create("myTrace.out")
	defer f.Close()
	_ = trace.Start(f)
	defer trace.Stop()

	//定义一个任务 //Context用于管理Goroutine之间的取消信号、请求作用域以及存储和传递上下文相关的值
	//使用这个根Context来创建一个新的任务，用于跟踪这个任务执行的过程
	//将跟踪信息与这个具体任务相关联
	ctx, task := trace.NewTask(context.Background(), "customerTask")
	defer task.End()
	var wg = sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		//启动十个协程，模拟做任务
		go func(num string) {
			defer wg.Done()

			trace.WithRegion(ctx, num, func() {
				var sum, i int64
				for ; i < 500000000; i++ {
					sum += 1
				}
				fmt.Println(num, sum)
			})

		}(fmt.Sprintf("num_%02d", i))
	}
}

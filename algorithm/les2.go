package main

import (
	"fmt"
)

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

// 输出什么       会报错   （对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的。）
//
//	func main() {
//		m["foo"].x = 4
//		fmt.Println(m["foo"].x)
//	}

// 用临时变量，或者定义map是value存指针地址。
//
//	func main() {
//		tmp := m["foo"]
//		tmp.x = 4
//		m["foo"] = tmp
//		fmt.Println(m["foo"].x)
//	}
const (
	mutexLocked = 1 << iota // 1左移0位，结果是1
	mutex1                  // 1左移1位，结果是2
	mutex2                  // 1左移2位，结果是4
	mutex3                  // 1左移3位，结果是8
)

func main() {
	fmt.Println(mutexLocked) // 输出: 1
	fmt.Println(mutex1)      // 输出: 2
	fmt.Println(mutex2)      // 输出: 4
	fmt.Println(mutex3)      // 输出: 8
}

//func main() {
//	var counter int64
//	var wg sync.WaitGroup
//	wg.Add(100)
//
//	for i := 0; i < 100; i++ {
//		go func() {
//			atomic.AddInt64(&counter, 2)
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//	fmt.Println("Counter:", atomic.LoadInt64(&counter))
//}

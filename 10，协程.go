package main

import (
	"fmt"
	"time"
)

// 一,协程在go语言中内部的表示,g结构体
// 1,g结构体中  stack 是debug中的栈信息    stack结构体中有  lo   hi    表示下限和上限 存储栈针信息
// 2, sched成员 ，gobuf结构体 目前程序运行现场
//
//	a,sp  原始指针  stackpointer
//	b,pc  程序计数器   运行到哪一行代码
//
// 3,atomicstatus 协程的状态
// 4,goid  协程的Id
// 5,m结构体 :描述操作系统线程的结构体
//
//			a,g0协程 ，操作调度器
//			b,curg 目前线程运行的哪个协程  g指针指向g结构体
//		    c,mOS  哪个操作系统线程信息
//	   单线程循环   g0为起始协程  1,schedule(proc.go)
//							  2,excute()
//						      3,gogo()
//						      4,业务方法
//							  5,goexit()
func main() {
	go go1()
	fmt.Println(time.Hour)
	time.Sleep(time.Hour)
}
func go1() {
	go2()
}

func go2() {
	go3()
}

func go3() {
	fmt.Println("go3")
}

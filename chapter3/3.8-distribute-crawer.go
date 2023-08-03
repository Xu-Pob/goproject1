package main

import (
	"fmt"
	"os"
	"unsafe"
)

// Master运行方式形如：go run 3.8-distribute-crawer.go master 127.0.0.1:8806
// 在这里 go run chapter3/3.8-distribute-crawer.go arg1 arg2....
// Worker运行方式形如：go run main.go 3.8-distribute-crawer.go 127.0.0.1:8806 127.0.0.1:8808 &
func main() {
	if len(os.Args) < 2 {
		sz := 100
		p := make([]int, sz)
		p[2] = 5
		fmt.Printf("%s: 使用日志请查看文件： %d\n", os.Args[0], *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + 2*unsafe.Sizeof(p[0]))))
	}
	switch os.Args[1] {
	case "master":
		fmt.Println("arg1:master")
	case "worker":
		fmt.Println("arg1:worker")
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	arr := []interface{}{1, 4, 6}
	fmt.Println(arr)
	fmt.Println(arr...)
	err := returnsError()
	if err != nil {
		fmt.Printf("操作失败: %v", err)
	}
}

// nil是一个合法的赋值给指针、接口、切片、映射、通道和函数类型变量的零值。
// 但是对于其他的类型声明，如结构体类型等，无法将nil直接赋值给变量
func returnsError() error {
	var err *os.PathError = nil
	fmt.Println(err)
	return err
}

package main

import "fmt"

// 类库  hmap 原理
// map的初始化
// 1，make     2,字面量
func main() {
	m := make(map[string]int, 10)
	fmt.Println(m)
}

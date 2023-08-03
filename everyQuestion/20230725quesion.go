package main

import "fmt"

// 输出什么
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//20,0,1,1
//2,0,1,1
//10,0,1,1
//1,,0,1,1

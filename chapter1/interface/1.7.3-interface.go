package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0)
	slice = append(slice, 1, 6, 8)
	var I interface{} = slice
	if res, ok := I.([]int); ok {
		fmt.Println(res)
		fmt.Println(ok)
	}

	Len(slice)

}

// 利用反射获取长度
//
//	func Len(array interface{}) int {
//		av := reflect.ValueOf(array)
//		return av.Len()
//	}

func Len(array interface{}) int {
	var length int
	if array == nil {
		length = 0
	}
	switch array.(type) {
	case []int:
		length = len(array.([]int))
	case []string:
		length = len(array.([]string))
	case []float32:
		length = len(array.([]float32))
	default:
		length = 0
	}
	fmt.Println(length)
	return length
}

package main

import (
	"fmt"
)

// 1，切片的本质是对数组的引用  slice    指针的引用，cap容量 ，len长度
// ,2，切片的创建
// a,根据数组创建 arr[0:3],slice[0:3]
// b,根据字面量创建
// c,make运行时创建数组
// 3，runtime包和reflect包有相同的数据结构
// 4，切片追加是会重建新的数组，并发上访问切片不安全
func main() {
	// b,根据字面量创建
	sl := []int{1, 2, 3} //==> j   ar:=[3]int{1,2,3}  slice{ar,3,3}
	// make运行时创建数组
	sl = make([]int, 10)
	//runtime.Slice
	//reflect.SliceHeader
	fmt.Println(sl)
}

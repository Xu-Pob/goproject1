package main

import (
	"fmt"
	"github.com/Xu-Pob/goproject1/chapter1/array/arrayLib"
	"sort"
)

func main() {
	a := []int{4, 3, 2, 1, 5, 6}
	//IntSlice 排序
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("排序后", a)

	b := [][]int{{1, 2, 3}, {1, 1, 1}, {54, 23, 2}, {1, 12, 12}}

	//二维数组排序
	result := arrayLib.ArraySort(b, 2-1)
	fmt.Println(result)

}

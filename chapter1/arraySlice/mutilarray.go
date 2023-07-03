package main

import (
	"fmt"
	"github.com/Xu-Pob/goproject1/chapter1/array/arrayLib"
)

func main() {
	//定义一个二维切片
	nums := [][]int{{1, 9, 5}, {2, 3, 6}, {3, 6, 9}, {1, 8, 3}}

	firstIndex := 2 //按第二列排序

	result := arrayLib.ArraySort(nums, firstIndex-1)

	fmt.Println(result)

	var array3 [1][2][3]int
	var array4 = [1][2][3]int{{{1, 2, 3}, {2, 1, 3}}}
	fmt.Println(array3)
	fmt.Println(array4)
	//var array5 = make([][],2)
	//fmt.Println(array5)
	c := make3D(2, 3, 4)
	c[1][0][0] = 9
	fmt.Println(c)
	print2D()
	defineShow2D()
}

// 三维数组生成器
func make3D(m, n, p int) [][][]float64 {
	buf := make([]float64, m*p*n)

	x := make([][][]float64, m)
	for i, _ := range x {
		x[i] = make([][]float64, n)
		for j := range x[i] {
			x[i][j] = buf[:p:p]
			buf = buf[p:]
		}
	}
	return x
}

////[[[0 0 0] [0 0 0]] [[0 0 9] [0 0 0]]]

func print2D() {
	// 二维数组（5 行 2 列
	var a = [5][2]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}}
	//输出数组元素
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

}

func defineShow2D() {
	var mArray [3][2]int
	mArray[0] = [2]int{5, 6}
	fmt.Println(mArray)
}

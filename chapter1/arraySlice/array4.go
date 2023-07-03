package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 每个go文件中，都可以有一个init函数，做一些初始化的工作，在main函数运行前，go框架会调用init函数
// go文件扫描顺序：全局变量定义->init函数->main函数
func init() {
	rand.New(rand.NewSource(time.Now().Unix()))
	//rand.Seed(time.Now().Unix())
}
func main() {
	strings := []string{"1", "2", "3", "4", "5", "6"}
	a, _ := RandomString(strings, 3)
	fmt.Println(a)
	fmt.Println(strings)
}
func findMaxValue() {
	array := []int{1, -2, 88, 66, 16, 68}
	maxValue := array[0]
	maxValueIndex := 0
	for i := 0; i < len(array); i++ {
		if maxValue < array[i] {
			maxValue = array[i]
			maxValueIndex = i
		}
	}
	fmt.Printf("maxValue=%v maxValueIndex=%v \n", maxValue, maxValueIndex)

}

// RandomString Fisher-Yates随机置乱算法生成随机整数
func RandomString(strings []string, length int) (string, error) {

	if len(strings) < 0 {
		return "", errors.New("字符长度不能小于0")
	}
	if length < 0 || length > len(strings) {
		return "", errors.New("参数长度非法")
	}

	for i := len(strings) - 1; i > 0; i-- {
		n := rand.Intn(i + 1)
		strings[i], strings[n] = strings[n], strings[i]
	}
	str := ""
	for _, s := range strings {
		str += s
	}
	return str, nil
}

// 删除重复的元素
func removeDuplicates(array []int) []int {

	//  1  1  4  1  1
	/// l:0 r:1  相同
	/// l:0  r:2  不同  a[0+1] = a[2]    ->  1 4 4 6 1
	/// l:1  r：3  不同  a[1+1] =a[3]     ->
	//
	return nil
}

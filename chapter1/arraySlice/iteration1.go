package main

import "fmt"

func main() {
	var nums = []int{1, 2, 4, 56, 6}
	for i, num := range nums {
		fmt.Printf("index:%d value:%d  element-address:%X \n", i, num, &nums[i])
	}
	fmt.Println("\n使用匿名变量（下划线）来忽略索引值：")
	//使用空白标识符（下划线）来忽略索引值
	for _, v := range nums {
		fmt.Printf(" value:%d   \n", v)
	}
	fmt.Println("\n使用 for 循环对切片进行迭代：")
	//使用传统的 for 循环对切片进行迭代
	for i := 0; i < len(nums); i++ {
		fmt.Printf("value:%d \n", nums[i])
	}
}

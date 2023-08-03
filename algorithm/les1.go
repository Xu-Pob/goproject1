package main

import "fmt"

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的
//那两个整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

//给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回[0,1]

func main() {
	result := fun3([]int{1, 2, 3, 5, 7}, 9)

	fmt.Println(result)
}

// 暴力法
func fun1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 两遍哈希
func fun2(nums []int, target int) []int {
	var temMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		a := target - nums[i]
		val, exists := temMap[a]
		if exists && val != nums[i] {
			return []int{i, val}
		}
	}
	return nil
}

// 一遍哈希
func fun3(nums []int, target int) []int {
	var temMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		a := target - nums[i]
		val, exists := temMap[a]
		if exists && val != nums[i] {
			return []int{i, val}
		}
		temMap[nums[i]] = i
	}
	return nil
}

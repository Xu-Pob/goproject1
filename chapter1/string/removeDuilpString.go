package main

import (
	"fmt"
)

func
//这种想法就是典型的贪心策略了：我们每次都找到当前能够移到最左边的、最小的字母。这就是我们得到结果的第一个字母，
//它之前的所有重复字母会被删掉；然后我们从它以后的字符串中，使用相同的逻辑，继续寻找第二个最小的字母。

main() {
	//println(len("abcadfsdf"))  //9
	removeString("abcadfsdf")
}

func removeString(s string) {
	runes := []rune(s)
	//var runesResult []rune
	left := 0
	for right := 1; right < len(runes); right++ {
		fmt.Print(runes[left], runes[right], "\n")

		if runes[left] > runes[right] {
			runes[left], runes[right] = runes[right], runes[left]
			//runesResult = append(runesResult, runes[left])
			left++
			continue
		}
		//left++
	}
	s = ""
	for _, v := range runes {
		s += string(v)
	}
	fmt.Println(s)
}

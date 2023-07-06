package main

import (
	"sort"
)

func main() {
	m := make(map[int]string)
	m[8] = "I"
	m[2] = "Love"
	m[6] = "go"
	sortedMap(m, func(k int, v string) {
		println(k, v)
	})
}

// 将map按照key排序
func sortedMap(m map[int]string, f func(k int, v string)) {
	var key []int
	for k, _ := range m {
		key = append(key, k)
	}
	sort.Ints(key)
	for _, k := range key {
		f(k, m[k])
	}
}

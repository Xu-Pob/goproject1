package main

import (
	"fmt"
	"sort"
)

type Programmera struct {
	Name string
	Age  int
}
type ProgrammerSlice []Programmera

func (s ProgrammerSlice) Len() int {
	return len(s)
}
func (s ProgrammerSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ProgrammerSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}
func main() {
	a := ProgrammerSlice{{
		Name: "Barry",
		Age:  30,
	},
		{
			Name: "Jack",
			Age:  22,
		},
		{
			Name: "Jim",
			Age:  18,
		},
	}
	//Stable是稳定的排序，相等时保留相对位置，，Sort是快速不稳定的排序
	sort.Sort(a)
	fmt.Println(a)
}

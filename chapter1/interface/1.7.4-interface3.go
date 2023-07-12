package main

import "fmt"

func main() {
	i := Minimum(1, 123, 456, 73, 9)
	j := Minimum(1.1, 66.6, 6.6, 8.5, 9.9, 16.6, -2.5)
	k := Minimum("e", "ab", "cd", "d")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}

//获取数组最小值，interface作为参数

func Minimum(first interface{}, rest ...interface{}) interface{} {
	minimum := first
	for _, v := range rest {
		switch v.(type) {
		case int:
			if v := v.(int); v < minimum.(int) {
				minimum = v
			}
		case float64:
			if v := v.(float64); v < minimum.(float64) {
				minimum = v
			}
		case string:
			if v := v.(string); v < minimum.(string) {
				minimum = v
			}
		}
	}
	return minimum
}

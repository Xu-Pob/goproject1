package main

import (
	"fmt"
	"reflect"
)

func main() {
	orgialstr := []string{"1", "张三", "李四"}
	result := ExsitString("2", orgialstr)

	fmt.Println("1 是否在 orgialstr 中", result)
	reflectVar()

	a := make([]string, 3)
	a[0] = "1"
	a[1] = "张三"
	index := arrayPosition(a, "1")
	fmt.Println(index)
}
func ExsitString(target string, array []string) bool {
	for _, s := range array {
		if target == s {
			return true
		}
	}
	return false
}
func reflectVar() {
	var num = 42
	value := reflect.ValueOf(num)

	fmt.Println("Type:", value.Type())       // 输出变量的类型
	fmt.Println("Kind:", value.Kind())       // 输出变量的种类
	fmt.Println("Value:", value.Int())       // 输出变量的值
	fmt.Println("Value:", value.Interface()) // 输出变量的值（使用空接口类型）

	// 修改变量的值（需要传入可寻址的值）
	pValue := reflect.ValueOf(&num)
	pValue.Elem().SetInt(99)
	fmt.Println("Modified Value:", num) // 输出变量的修改后的值
}

// 查找一个元素在切片的位置
func arrayPosition(arr interface{}, d interface{}) int {
	array := reflect.ValueOf(arr)
	for i := 0; i < array.Len(); i++ {
		v := array.Index(i)
		if v.Interface() == d {
			return i
		}
	}
	return -1
}

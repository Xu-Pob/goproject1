package main

import (
	"fmt"
	"unicode/utf8"
)

// rune 类型（实际上就是 int32 类型）来表示 Unicode 码。rune 字面值格式是用单引号包围起来的一个字符
func main() {
	//聲明一個字符串變量str
	str := "在go中可以通過切片截取一個數組或字符串"
	fmt.Println(utf8.RuneCountInString(str))
	//打印字节长度
	fmt.Println(len(str))
	//子字符串操作 s[i:j] 产生一个新的字符串，它包含原始字符串从索引 i 开始到索引 j-1 的字节
	str1 := str[0:3]
	fmt.Println(str1)
	nameRune := []rune(str)
	fmt.Println(len(nameRune))
	//截取后的字符串前10个字符
	fmt.Println(string(nameRune[:10]))
}

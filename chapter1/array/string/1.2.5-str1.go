package main

import (
	"fmt"
	"regexp"
	"strings"
)

// strings.   字符串處理的包
func main() {
	a := "张三丰"
	b := &a
	fmt.Println(b)

	fmt.Println(strings.ToUpper("helloworld"))
	fmt.Println(strings.ToLower("helWERWERrld"))

	//stringSplitTest()
	stringTrim()
}

// 字符串分割,正则表达式
func stringSplitTest() {
	test := "Hi,You,Love"
	str := test
	keywordslice := strings.Split(test, ",")
	for _, v := range keywordslice {
		reg := regexp.MustCompile("(?i)" + v)
		str = reg.ReplaceAllString(str, strings.ToUpper(v))
		println(str)
	}
}

func stringTrim() {
	str := "   a b "
	str1 := strings.TrimSpace(str)
	str2 := strings.Trim(str, " ")
	println(str1)
	println(str2)
}

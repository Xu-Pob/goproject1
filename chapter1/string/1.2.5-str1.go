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

	stringSplitTest()
	stringTrim()
	SpliceByRune()
	JoinString()
	SplitString()
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

func SpliceByRune() {
	fn := func(c rune) bool {
		return strings.ContainsRune(",|/", c)
	}
	str := strings.FieldsFunc("Python,Jquery|JavaScript,Go,C++/C", fn)
	fmt.Println(str)
	//[Python Jquery JavaScript Go C++ C]

	str1 := strings.TrimFunc("|Python,Jquery|JavaScript,Go,C++/C", fn)
	fmt.Println(str1)
}

func JoinString() {
	str := []string{"I", "Love", "Go", "Java"}
	res := strings.Join(str, "-")
	fmt.Println(res)
}
func SplitString() {
	s := "I_Love_Go_Advanced"
	res1 := strings.SplitN(s, "_", 2)
	for i := range res1 {
		fmt.Println(res1[i])
	}

	//切割后包含分隔符
	res2 := strings.SplitAfter(s, "_")
	for i := range res2 {
		fmt.Println(res2[i])
	}
}

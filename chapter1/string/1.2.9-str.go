package main

import "regexp"

// 匹配中文字符
var cnRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

// 指针，引用传递，得值后赋值地址，   值传递，只是改变值，main中的str的值还是没变
func main() {
	str := "我是a,你是b"
	StrFilterGetChinese(&str)
	println(str)
}

//匹配中文字符并返回

func StrFilterGetChinese(src *string) {
	strNew := ""
	for _, c := range *src {
		if cnRegexp.MatchString(string(c)) {
			strNew += string(c)
		}
	}
	*src = strNew
}

//我是你是

//

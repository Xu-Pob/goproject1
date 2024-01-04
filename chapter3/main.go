package main

import "github.com/Xu-Pob/goproject1/chapter3/crawler"

func main() {
	crawler.Do([]string{
		//"https://baijiahao.baidu.com/s?id=1771362463174303172&wfr=spider&for=pc",
		//路径加上https请求
		"https://www.baidu.com",
	})
}

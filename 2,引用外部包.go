package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
)

// 如果报红色字体，检查一下goproxy是否设置正确
// go get 外部包地址【@版本号】（在终端中输入命令)
func main() {
	c := tunny.Pool{}
	fmt.Println(&c)
}

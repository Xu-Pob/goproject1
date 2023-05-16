package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
)

// go get 外部包地址【@版本号】（在终端中输入命令)
func main() {
	c := tunny.Pool{}
	fmt.Println(&c)
}

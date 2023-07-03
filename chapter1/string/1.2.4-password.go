package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	fmt.Println(CreatePassword())
}
func CreatePassword() string {
	//t := time.Now()
	h := md5.New()
	//io.WriteString(h, "shirdon.liao")
	//io.WriteString(h, t.String())
	password := fmt.Sprintf("%x", h.Sum(nil))
	return password
}

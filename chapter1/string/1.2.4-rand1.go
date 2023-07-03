package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	b := getRandomString(22)
	fmt.Println(b)
}

// 獲取隨機字符串
func getRandomString(l int) string {

	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	//rand.NewSource(time.Now().UnixNano())

	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	println(r.Int31())
	println(r.Intn(len(bytes)))
	return string(result)
}

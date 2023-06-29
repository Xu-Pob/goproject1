package main

import (
	"fmt"
	"math/rand"
)

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func main() {
	b := make([]byte, 4)
	fmt.Println(len(StdChars))
	fmt.Println(b)
	fmt.Println(6 + 6/4)
	for i := 0; i < 4; i++ {
		if _, err := rand.Read(b); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}

	}
	fmt.Println(b)
}

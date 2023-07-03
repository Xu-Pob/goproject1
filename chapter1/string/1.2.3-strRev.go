package main

import "fmt"

func main() {
	str := "12323343"
	strRev := Reversal(str)
	fmt.Println(strRev)
}

func Reversal(a string) (re string) {
	//將字符串轉為rune數組
	b := []rune(a)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	re = string(b)
	return
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	// 初始化一个切片 seq
	seq := []string{"i", "love", "go", "advanced", "programming"}
	// 指定删除位置
	index := 2
	seq, err := deleteIndex(seq, index)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(seq)
	//strings.join
	fmt.Println(arrayToString(seq))
}
func deleteIndex(strings []string, index int) ([]string, error) {

	if index < 0 || index > len(strings) {
		return nil, errors.New("参数长度错误")
	}
	//连接
	return append(strings[:index], strings[index+1:]...), nil
}

func arrayToString(arrays []string) string {
	str := ""
	for _, v := range arrays {
		str += v
	}
	return str
}

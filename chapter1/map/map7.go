package main

import "fmt"

func main() {
	// 创建map
	m := map[string]string{
		"a": "value_a",
		"b": "value_b",
	}
	var sli []*string
	for k, v := range m {
		//k一直使用同一块内存，v也一样
		fmt.Printf("k[%v]:[%p].v[%v]:[%p]\n", k, &k, v, &v)
		sli = append(sli, &v)
	}
	// 输出
	//k:[0xc000010200].v:[0xc000010210]
	//k:[0xc000010200].v:[0xc000010210]
	for _, s := range sli {
		fmt.Println(*s)
	}
	// 输出
	// value_b
	// value_b
}

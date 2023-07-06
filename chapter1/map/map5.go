package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type People struct {
	Name string
	Age  int
}

func main() {
	ps := []People{
		{Name: "张三", Age: 18},
		{Name: "张四", Age: 16},
		{Name: "张五", Age: 19},
	}
	//根据什么排序
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Age < ps[j].Age || ps[i].Name < ps[j].Name
	})
	//将结构体数组转为json
	jsonData, err := json.Marshal(ps)
	if err != nil {
		fmt.Println(err)
	}
	jsonStr := string(jsonData)

	//将json转为map
	psmap := make([]map[string]interface{}, len(ps))
	err = json.Unmarshal(jsonData, &psmap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ps)
	fmt.Println(jsonStr)
	fmt.Println(psmap)

	for i, p := range psmap {
		fmt.Printf("第%d号：", i)
		for k, v := range p {
			fmt.Printf("k:[%v].v:[%v]  ", k, v)
		}
		fmt.Printf("\n")
	}
}

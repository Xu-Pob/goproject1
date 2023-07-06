package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	sortmap()
}

func GetMapKey() {
	//声明并初始化一个map，key是int64类型，value是string类型
	varMap := make(map[int64]string)
	varMap[1] = "Go"
	varMap[2] = "Advanced"

	for _, num := range []int64{1, 2, 3, 4} {
		if v, ok := varMap[num]; ok {
			fmt.Printf("varMap中包含key:%d,value值为:%s\n", num, v)
		} else {
			fmt.Printf("varMap中不包含key:%d,value值为:%s\n", num, v)
		}
	}
}

func MapToJson() {
	instance := map[string]interface{}{
		"name":   "Pobo",
		"goodAt": "Go Programming",
	}
	//转为json字符串的字节数组
	jsonStr, err := json.Marshal(instance)
	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}

	fmt.Println(string(jsonStr))
}

func JsonToMap() {
	jsonStr := `{
			"name": "Pobo",
			"goodAt":"GO programming"
			}`
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mapResult)
}

func sortmap() {
	arr := make(map[int]int, 5)
	arr[0] = 88
	arr[1] = 66
	arr[2] = 99
	b := []int{}
	fmt.Println("排序前")
	for k, v := range arr {
		fmt.Println(k, v)
		b = append(b, v)
	}
	fmt.Println("排序后")
	sort.Ints(b)
	for k, v := range b {
		fmt.Println(k, v)
	}
}

//排序前的值如下：
//0 88
//1 66
//2 99
//排序后的值如下：
//0 66
//1 88
//2 99

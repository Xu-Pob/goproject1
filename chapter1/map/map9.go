package main

import (
	"fmt"
	"sort"
)

func main() {
	f := frequencies("adsdsad")
	fmt.Println(f)
}
func mapSort() {
	counts := map[string]int{"Barry": 96, "Aaron": 98, "Clan": 97}
	//fmt.Println(len(counts))

	keys := make([]string, 0, len(counts))
	for k, _ := range counts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return counts[keys[i]] < counts[keys[j]]
	})
	for _, key := range keys {
		fmt.Printf("%s, %d\n", key, counts[key])
	}
}

//Barry, 96
//Clan, 97
//Aaron, 98

// 计算字符串的字符出现频率，并按照频率降序排序
func frequencies(s string) []frequency {

	mapfqs := make(map[string]int)
	frequencys := make([]frequency, 0, len(s))
	for _, v := range s {
		mapfqs[string(v)]++
	}
	for k, v := range mapfqs {
		frequencys = append(frequencys, frequency{char: k, fre: v})
	}

	sort.Slice(frequencys, func(i, j int) bool {
		return frequencys[i].fre > frequencys[j].fre
	})
	//fmt.Println(mapfqs)
	return frequencys
}

type frequency struct {
	char string
	fre  int
}

package main

import "fmt"

func main() {
	mapInterface()
}
func baseMap() {
	m := map[string]map[string]string{}
	m["programmer"] = map[string]string{}
	m["programmer"]["name"] = "Pobo"
	m["manager"] = map[string]string{}
	m["manager"]["goodAt"] = "Go"
	fmt.Println(m["programmer"]["name"])
	fmt.Println(m)

	for key, value := range m {
		for k, v := range value {
			fmt.Println("k:", k, "V:", v)
		}
		fmt.Println("k:", key, "V:", value)
	}
}

//Pobo
//map[manager:map[goodAt:Go] programmer:map[name:Pobo]]
//k: goodAt V: Go
//k: manager V: map[goodAt:Go]
//k: name V: Pobo
//k: programmer V: map[name:Pobo]

func mapInterface() {
	mainMap := map[string]interface{}{}
	subMap := make(map[string]string)
	subMap["Key_1"] = "SubValue_1"
	subMap["Key_2"] = "SubValue_2"
	mainMap["subMap"] = subMap
	for key, val := range mainMap {
		for subKey, subVal := range val.(map[string]string) {
			fmt.Printf("mapname:=%s  Key=%s  value=%s\n", key, subKey, subVal)
		}
	}
}

//mapname:=subMap  Key=Key_1  value=SubValue_1
//mapname:=subMap  Key=Key_2  value=SubValue_2

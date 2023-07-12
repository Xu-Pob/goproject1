package main

import (
	"fmt"
	"reflect"
)

type Programmer struct {
	Name string `json:"name" level:"12"`
}

func main() {
	pro := Programmer{}
	//反射获取标签信息
	typeofPro := reflect.ValueOf(pro)
	PType := typeofPro.Type() //typeOfPro := reflect.TypeOf(pro)

	proFieldName, ok := PType.FieldByName("Name")
	if ok {
		//打印标签信息
		fmt.Println(proFieldName.Tag.Get("json"), proFieldName.Tag.Get("level"))
	}

}

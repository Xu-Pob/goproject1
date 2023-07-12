package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string
	GoodAt string
}

func main() {

	user := User{"Pobo", "Go"}
	//匿名结构体
	data := struct {
		Title string
		User  User
	}{
		"info",
		user,
	}
	fmt.Println(data)
	showResponse()
}

type Item struct {
	Title string
	URL   string
}
type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func showResponse() {
	jsonStr := `{
	"data": {
		"children": [
			{
				"data": {
					"title": "Shirdon's Blog'",
					"url": "https://www.shirdon.com"
				}
			}
		]
	}
}`
	res := Response{}
	json.Unmarshal([]byte(jsonStr), &res)
	fmt.Println(res)
}

package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()
	//赋值
	_, err = conn.Do("SET", "Pobo", 18)
	if err != nil {
		log.Fatal("redis set error:", err)
	}

	name, err := redis.String(conn.Do("GET", "Pobo"))

	if err != nil {
		log.Fatal("redis get error:", err)
	} else {
		fmt.Printf("Get name: %s \n", name)
	}
}

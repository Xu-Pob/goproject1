package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal("connect redis server:", err)
	}
	fmt.Println(conn)
	defer conn.Close()
}

package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()
	//集合中增加用户分数
	_, err = conn.Do("ZADD", "score", 99, "Jack")
	if err != nil {
		panic(err)
	}
	_, err = conn.Do("ZADD", "score", 97, "James", 85, "Pobo")
	if err != nil {
		panic(err)
	}
	//获取成员个数
	result, err := conn.Do("ZCARD", "score")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	//取出并升序
	scoremap, err := redis.StringMap(conn.Do("ZREVRANGE", "score", 0, 2, "withscores"))
	for name := range scoremap {
		fmt.Println(name, scoremap[name])
	}
	//取出并降序
	scoremap, err = redis.StringMap(conn.Do("ZRANGE", "score", 0, 2, "withscores"))
	for name := range scoremap {
		fmt.Println(name, scoremap[name])
	}
	//取出POBO的分数
	score, err := redis.Int64(conn.Do("ZSCORE", "score", "Pobo"))

	if err != nil {
		panic(err)
	}
	fmt.Println(score)
	//移除集合中的某一个或者多个成员
	result, err = conn.Do("ZREM", "score", "Pobo")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

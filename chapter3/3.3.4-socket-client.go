package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//	func Log(msg string, err error) {
//		fmt.Println(msg, err)
//		os.Exit(1)
//	}
func main() {
	//拨号
	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		Log("拨号失败", err)
	}
	defer conn.Close()
	for {
		//发送数据
		//读取输入行
		reader := bufio.NewReader(os.Stdin)
		readString, err := reader.ReadString('\n')
		if err != nil {
			Log("读取输入数据失败", err)
		}
		conn.Write([]byte(readString))

		//读取数据，并显示
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			Log("读取监听数据失败", err)
		}
		fmt.Printf("%v说:%v", conn.RemoteAddr(), string(buf[:n]))
	}

}

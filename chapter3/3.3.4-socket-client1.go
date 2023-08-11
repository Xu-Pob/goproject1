package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//	func Log(msg string, err error) {
//		fmt.Println(msg, err)
//		os.Exit(1)
//	}
func main() {
	//拨号
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8888,
	})
	if err != nil {
		fmt.Println("拨号失败：", err)
	}
	defer conn.Close()
	for {
		//发送数据
		//读取输入行
		reader := bufio.NewReader(os.Stdin)
		readString, _ := reader.ReadString('\n')
		if strings.ToUpper(readString) == "Q" {
			fmt.Println("退出聊天系统～")
			return
		}
		conn.Write([]byte(readString))

		//读取数据，并显示
		var buf [1024]byte
		n, _, _ := conn.ReadFromUDP(buf[:])

		fmt.Printf("%v说:%v", conn.RemoteAddr(), string(buf[:n]))
	}

}

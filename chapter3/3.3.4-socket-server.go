package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Log(msg string, err error) {
	fmt.Println(msg, err)
	os.Exit(1)
}
func Work(sock net.Conn) {
	defer sock.Close()
	//读取数据
	for {
		//数据流读取 ,接受客户端的请求
		var buf [1024]byte
		n, err := sock.Read(buf[:])
		if err != nil {
			Log("读取数据出错", err)
		}
		//显示客户端数据
		fmt.Printf("%v说: %v", sock.RemoteAddr(), string(buf[:n]))

		//创建写入器， 写入读取的信息并发送
		reader := bufio.NewReader(os.Stdin)
		readString, err := reader.ReadString('\n')
		if err != nil {
			Log("读取输入的失败", err)
		}

		_, err = sock.Write([]byte(readString))
		if err != nil {
			Log("发送失败", err)
		}
	}
}

func main() {

	//监听
	listen, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		Log("监听失败", err)
	}
	defer listen.Close()
	//接受请求处理
	for {
		conn, err := listen.Accept()
		if err != nil {
			Log("接受请求失败", err)

		}
		go Work(conn)
	}

}

package main

import (
	"context"
	"fmt"
	pb "github.com/Xu-Pob/goproject1/chapter3/protobuf"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//创建一个不安全的链接
	conn, err := grpc.Dial(":8069", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("dial error: %v\n", err)

	}
	defer conn.Close()

	//实例化
	client := pb.NewStudentServiceClient(conn)

	//调用服务
	req := new(pb.Request)
	req.Name = "barry"
	resp, err := client.GetStudentInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}
	fmt.Printf("Recevied: %v\n", resp)

}

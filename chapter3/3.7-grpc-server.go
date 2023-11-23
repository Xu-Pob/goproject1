package main

import (
	"context"
	"fmt"
	pb "github.com/Xu-Pob/goproject1/chapter3/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

// StudentServiceServer 定义服务结构体
type StudentServiceServer struct {
	pb.UnimplementedStudentServiceServer
}

func (p *StudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}
func (p *StudentServiceServer) GetStudentInfo(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "barry" {
		resp = &pb.Response{
			Uid:      8,
			Username: name,
			Grade:    "6",
			GoodAt:   []string{"语文", "数学", "计算机"},
		}
	}
	err = nil
	return
}
func main() {
	port := ":8069"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n", port)
	s := grpc.NewServer()
	pb.RegisterStudentServiceServer(s, &StudentServiceServer{})
	s.Serve(l)
}

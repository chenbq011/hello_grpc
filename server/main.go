package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"hello_grpc/hello/pb"
	"net"
)

type server struct {

}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message:"Hello " + in.Name},nil
}


func main()  {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("failed to listen: %v",err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

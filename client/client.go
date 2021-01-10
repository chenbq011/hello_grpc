package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello_grpc/hello/pb"
)

func main()  {
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name:"Bob"})
	if err != nil {
		fmt.Printf("failed to greet: %v", err)
	}
	fmt.Printf("Greeting: %s \n", r.Message)
}

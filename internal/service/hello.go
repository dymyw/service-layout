package service

import (
	"context"
	"log"

	pb "github.com/dymyw/service-layout/api/hello/v1"
)

type HelloServer struct {
	pb.UnimplementedGreeterServer
}

func (hs *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("HelloServer Received: %v", in.GetName())

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

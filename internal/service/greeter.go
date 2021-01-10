package service

import (
	"context"
	"log"

	pb "github.com/dymyw/service-layout/api/hello/v1"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

func (hs *GreeterService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("GreeterService Received: %v", in.GetName())

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

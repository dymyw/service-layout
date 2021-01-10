package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dymyw/service-layout/api/account/v1"
	"github.com/dymyw/service-layout/api/hello/v1"
	"github.com/dymyw/service-layout/internal/service"
)

const port = ":50001"

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &service.GreeterService{})
	account.RegisterAccountServer(s, &service.AccountService{})

	// for test
	reflection.Register(s)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

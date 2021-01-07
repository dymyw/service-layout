package main

import (
	"github.com/dymyw/service-layout/api/account/v1"
	"log"
	"net"

	"google.golang.org/grpc"

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
	hello.RegisterGreeterServer(s, &service.HelloServer{})
	account.RegisterAccountServiceServer(s, &service.AccountServer{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/dymyw/service-layout/api/account/v1"
	"github.com/dymyw/service-layout/api/hello/v1"
	"github.com/dymyw/service-layout/internal/service"
)

const port = ":50001"

func main() {
	log.Println("server start")

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("start grpc listen addr: [::]::" + port)

	server := grpc.NewServer()
	hello.RegisterGreeterServer(server, &service.GreeterService{})
	account.RegisterAccountServer(server, &service.AccountService{})

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

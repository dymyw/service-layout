package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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

	server := grpc.NewServer()
	hello.RegisterGreeterServer(server, &service.GreeterService{})
	account.RegisterAccountServer(server, &service.AccountService{})

	exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
	c := make(chan os.Signal, len(exitSignals))
	signal.Notify(c, exitSignals...)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())

		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			server.Stop()
			log.Println("server stop")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

	// for test
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

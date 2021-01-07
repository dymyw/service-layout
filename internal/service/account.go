package service

import (
	"context"
	"log"

	"github.com/dymyw/service-layout/internal/biz"
	"github.com/dymyw/service-layout/internal/data"

	pb "github.com/dymyw/service-layout/api/account/v1"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
}

func (as *AccountServer) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInReply, error) {
	log.Printf("AccountServer Received: %v", in.GetName())

	// DTO -> DO
	account := new(biz.Account)
	account.Name = in.GetName()

	// repo
	accountRepo := data.NewAccountRepo()
	// user case
	accountUserCase := biz.NewAccountUserCase(accountRepo)
	// save
	result := accountUserCase.SaveAccount(account)

	return &pb.SignInReply{Result: result}, nil
}

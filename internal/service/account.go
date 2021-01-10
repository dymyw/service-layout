package service

import (
	"context"
	"log"

	"github.com/dymyw/service-layout/internal/biz"
	"github.com/dymyw/service-layout/internal/data"

	pb "github.com/dymyw/service-layout/api/account/v1"
)

type AccountService struct {
	pb.UnimplementedAccountServer
}

func (as *AccountService) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInReply, error) {
	log.Printf("AccountService SignIn Received: %v", in.GetName())

	name := in.GetName()

	// DTO -> DO
	account := new(biz.Account)
	account.Name = name

	// db
	db, _, _ := data.NewDB()

	// repo
	accountRepo := data.NewAccountRepo(db)
	// user case
	accountUserCase := biz.NewAccountUserCase(accountRepo)
	// save
	result := accountUserCase.SaveAccount(account)

	return &pb.SignInReply{Result: result}, nil
}

func (as *AccountService) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	log.Printf("AccountService GetAccount Received: %v", in.GetId())

	id := in.GetId()

	// db
	db, _, _ := data.NewDB()

	// repo
	accountRepo := data.NewAccountRepo(db)
	// user case
	accountUserCase := biz.NewAccountUserCase(accountRepo)
	// account
	account, err := accountUserCase.GetInfo(id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAccountReply{Name: account.Name}, nil
}

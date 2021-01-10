package service

import (
	"context"
	"errors"
	"log"

	"github.com/dymyw/service-layout/internal/biz"
	"github.com/dymyw/service-layout/internal/data"

	pb "github.com/dymyw/service-layout/api/account/v1"
)

var UserNotFound = errors.New("user not found")

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
}

func (as *AccountServer) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInReply, error) {
	log.Printf("AccountServer SignIn Received: %v", in.GetName())

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

func (as *AccountServer) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	log.Printf("AccountServer GetAccount Received: %v", in.GetId())

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
		if errors.Is(err, biz.NotFound) {
			return nil, UserNotFound
		}
	}

	return &pb.GetAccountReply{Name: account.Name}, nil
}

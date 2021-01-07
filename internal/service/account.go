package service

import (
	"github.com/dymyw/service-layout/internal/biz"
	"github.com/dymyw/service-layout/internal/data"
)

func SignIn(name string) {
	// DTO -> DO
	account := new(biz.Account)
	account.Name = name

	// repo
	accountRepo := data.NewAccountRepo()
	// user case
	accountUserCase := biz.NewAccountUserCase(accountRepo)
	// save
	accountUserCase.SaveAccount(account)
}

package data

import (
	"fmt"
	"github.com/dymyw/service-layout/internal/biz"
)

// NewAccountRepo
func NewAccountRepo() biz.AccountRepo {
	return &accountRepo{}
}

// accountRepo
type accountRepo struct {}

// SaveAccount
func (ar *accountRepo) SaveAccount(a *biz.Account) bool {
	fmt.Println(a)

	return true
}

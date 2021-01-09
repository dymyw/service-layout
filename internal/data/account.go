package data

import (
	"database/sql"
	"fmt"
	"github.com/dymyw/service-layout/internal/biz"
)

// NewAccountRepo
func NewAccountRepo(db *sql.DB) biz.AccountRepo {
	return &accountRepo{db: db}
}

// accountRepo
type accountRepo struct {
	db *sql.DB
}

// SaveAccount
func (ar *accountRepo) SaveAccount(a *biz.Account) bool {
	fmt.Println(a)

	return true
}

// GetInfo
func (ar *accountRepo) GetInfo(id string) biz.Account {
	var mobile string
	err := ar.db.QueryRow("SELECT mobile FROM t_user where id = ?", id).Scan(&mobile)
	if err != nil {
		panic(err)
	}

	return biz.Account{Name: mobile}
}

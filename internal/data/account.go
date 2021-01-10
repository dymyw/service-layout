package data

import (
	"database/sql"
	"fmt"
	"github.com/dymyw/service-layout/internal/biz"
)

var _ biz.AccountRepo = &accountRepo{}
//var _ biz.AccountRepo = new(accountRepo)
//var _ biz.AccountRepo = (*accountRepo)(nil)

// todo wire error
//var AccountRepoSet = wire.NewSet(
//	NewAccountRepo,
//	// bind interface and struct
//	wire.Bind(new(biz.AccountRepo), new(*accountRepo)),
//)

// NewAccountRepo
func NewAccountRepo(db *sql.DB) *accountRepo {
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

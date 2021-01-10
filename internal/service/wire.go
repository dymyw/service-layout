//+build wireinject

package service

import (
	"database/sql"
	"github.com/dymyw/service-layout/internal/biz"
	"github.com/dymyw/service-layout/internal/data"
	"github.com/google/wire"
)

// todo wire error
func InitializeAccountUserCase(db *sql.DB) biz.AccountUserCase {
	wire.Build(biz.NewAccountUserCase, data.NewAccountRepo)

	return biz.AccountUserCase{}
}

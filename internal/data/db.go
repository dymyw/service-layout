package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (db *sql.DB, cf func(), err error) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/common")
	if err != nil {
		panic(err)
	}

	cf = func() {db.Close()}
	return
}

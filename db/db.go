package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:新しいパスワード@/financial_app")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// transactions.go
package models

import (
	"database/sql"
	"fmt"
)

func InsertTransaction(db *sql.DB, transType string, amount float64, date string) error {
	insertStatement := "INSERT INTO transactions (type, amount, date) VALUES (?, ?, ?)"
	_, err := db.Exec(insertStatement, transType, amount, date)
	return err
}

func GetTransactions(db *sql.DB) {
	query := "SELECT id, type, amount, date FROM transactions"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// 取得したデータを表示
	for rows.Next() {
		var id int
		var transType string
		var amount float64
		var date string
		if err := rows.Scan(&id, &transType, &amount, &date); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Type: %s, Amount: %.2f, Date: %s\n", id, transType, amount, date)
	}
}

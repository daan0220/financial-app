package models

import (
	"database/sql"
)

// Transaction データモデル
type Transaction struct {
	ID     int
	Type   string
	Amount float64
	Date   string
}

// InsertTransaction は新しい取引をデータベースに追加
func InsertTransaction(db *sql.DB, transType string, amount float64, date string) error {
	insertStatement := "INSERT INTO transactions (type, amount, date) VALUES (?, ?, ?)"
	_, err := db.Exec(insertStatement, transType, amount, date)
	return err
}

// GetAllTransactions はすべての取引を取得
func GetAllTransactions(db *sql.DB) ([]Transaction, error) {
	query := "SELECT id, type, amount, date FROM transactions"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.ID, &t.Type, &t.Amount, &t.Date); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

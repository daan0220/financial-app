package handlers

import (
	"encoding/json"
	"financial_app/db"
	"financial_app/models"
	"net/http"
)

// GetTransactions はすべての取引を取得するハンドラー
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	// データベースから取引を取得
	db, err := db.NewDB() // データベース接続を確立
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	transactions, err := models.GetAllTransactions(db) // データベース接続を渡して取引を取得
	if err != nil {
		http.Error(w, "Failed to get transactions", http.StatusInternalServerError)
		return
	}

	// JSON形式で取引をクライアントに返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

// CreateTransaction は新しい取引を作成するハンドラー
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// リクエストからデータをパース
	var newTransaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&newTransaction)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	db, err := db.NewDB() // データベース接続を確立
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// 取引をデータベースに追加
	err = models.InsertTransaction(db, newTransaction.Type, newTransaction.Amount, newTransaction.Date) // データベース接続と新しい取引データを渡して取引を追加
	if err != nil {
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	// 成功したことをクライアントに通知
	w.WriteHeader(http.StatusCreated)
}

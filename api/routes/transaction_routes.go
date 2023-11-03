package routes

import (
	"financial_app/api/handlers"
	"net/http"
)

func SetupRoutes() {
	// ルートハンドラーを設定
	http.HandleFunc("/api/transactions", handlers.GetTransactions)
	http.HandleFunc("/api/transactions/create", handlers.CreateTransaction)
}

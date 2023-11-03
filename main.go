package main

import (
	"financial_app/api/routes"
	"financial_app/db"
	"fmt" // fmtパッケージをインポート
	"net/http"
)

func main() {
	fmt.Println("Connecting to the database...")
	db, err := db.NewDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		panic(err)
	}
	defer db.Close()

	// APIエンドポイントの設定
	fmt.Println("Setting up API routes...")
	routes.SetupRoutes()
	routes.SetupUserRoutes()
	fmt.Println("API routes are set up.")

	fmt.Println("Starting the server on port 8080...")
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running.")

}

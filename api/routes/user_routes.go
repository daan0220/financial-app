package routes

import (
	"financial_app/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupUserRoutes() {
	// ユーザー関連のルートを設定
	r := mux.NewRouter()

	r.HandleFunc("/api/users", handlers.GetUser).Methods("GET")
	r.HandleFunc("/api/users/create", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/update/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/delete/{id}", handlers.DeleteUser).Methods("DELETE")

	http.Handle("/", r)
}

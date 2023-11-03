package handlers

import (
	"encoding/json"
	"financial_app/db"
	"financial_app/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUser は特定のユーザーを取得するハンドラー
func GetUser(w http.ResponseWriter, r *http.Request) {
	// データベースから取引を取得
	db, err := db.NewDB() // データベース接続を確立
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// リクエストからユーザーIDを取得
	userIDD := r.URL.Query().Get("id") // ユーザーIDを文字列として取得

	userID, err := strconv.Atoi(userIDD) // 文字列をintに変換
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// データベースからユーザーを取得
	user, err := models.GetUser(db, userID) // データベース接続とユーザーIDを渡してユーザーを取得
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// JSON形式でユーザーをクライアントに返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser は新しいユーザーを作成するハンドラー
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// リクエストからデータをパース
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
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

	// ユーザーをデータベースに追加
	err = models.InsertUser(db, newUser.Username, newUser.Email) // データベース接続と新しいユーザーデータを渡してユーザーを追加
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// 成功したことをクライアントに通知
	w.WriteHeader(http.StatusCreated)
}

// UpdateUser はユーザー情報を更新するハンドラー
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// パスパラメータからユーザーIDを取得
	vars := mux.Vars(r)
	userIDStr := vars["id"]

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// リクエストから新しいユーザー情報を取得
	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	db, err := db.NewDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// データベースからユーザー情報を更新
	err = models.UpdateUser(db, userID, updatedUser)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// 成功したことをクライアントに通知
	w.WriteHeader(http.StatusOK)
}

// DeleteUser はユーザーを削除するハンドラー
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// パスパラメータからユーザーIDを取得
	vars := mux.Vars(r)
	userIDStr := vars["id"]

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db, err := db.NewDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// データベースからユーザー情報を削除
	err = models.DeleteUser(db, userID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	// 成功したことをクライアントに通知
	w.WriteHeader(http.StatusOK)
}

package models

import (
	"database/sql"
)

// User データモデル
type User struct {
	ID       int
	Username string
	Email    string
}

// InsertUser は新しいユーザーをデータベースに追加
func InsertUser(db *sql.DB, username string, email string) error {
	insertStatement := "INSERT INTO users (username, email) VALUES (?, ?)"
	_, err := db.Exec(insertStatement, username, email)
	return err
}

// GetUser は指定されたユーザーを取得
func GetUser(db *sql.DB, userID int) (*User, error) {
	query := "SELECT id, username, email FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)

	var u User
	err := row.Scan(&u.ID, &u.Username, &u.Email)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UpdateUser は指定されたユーザーの情報を更新
func UpdateUser(db *sql.DB, userID int, updatedUser User) error {
	updateStatement := "UPDATE users SET username = ?, email = ? WHERE id = ?"
	_, err := db.Exec(updateStatement, updatedUser.Username, updatedUser.Email, userID)
	return err
}

// DeleteUser は指定されたユーザーを削除
func DeleteUser(db *sql.DB, userID int) error {
	deleteStatement := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(deleteStatement, userID)
	return err
}

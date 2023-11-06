package config

import (
	"os"
	"strconv"
)

// AppConfig はアプリケーションの設定情報を管理
type AppConfig struct {
	DatabaseURL string
	APIPort     int
}

// LoadConfig は設定情報を読み込み
func LoadConfig() *AppConfig {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		// 環境変数から設定を読み込めない場合のデフォルト値
		databaseURL = "mysql://root:パスワード@/financial_app"
	}

	apiPortStr := os.Getenv("API_PORT")
	var apiPort int
	if apiPortStr != "" {
		var err error
		apiPort, err = strconv.Atoi(apiPortStr)
		if err != nil {
			// パースエラーが発生した場合のデフォルト値
			apiPort = 8080
		}
	} else {
		// 環境変数が設定されていない場合のデフォルト値
		apiPort = 8080
	}

	return &AppConfig{
		DatabaseURL: databaseURL,
		APIPort:     apiPort,
	}
}

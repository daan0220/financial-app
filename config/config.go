package config

// AppConfig はアプリケーションの設定情報を管理
type AppConfig struct {
	DatabaseURL string
	APIPort     int
}

// LoadConfig は設定情報を読み込み
func LoadConfig() *AppConfig {
	// 実際の設定情報をここで読み込む
	return &AppConfig{
		DatabaseURL: "mysql://root:新しいパスワード@/financial_app",
		APIPort:     8080,
	}
}

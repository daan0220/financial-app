# ベースイメージとしてGo 1.18イメージを使用
FROM golang:1.18

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールを初期化
COPY go.mod go.sum ./
RUN go mod download

# プロジェクトのソースコードをコピー
COPY . .

# Goアプリケーションをビルド
RUN go build -o main

# ポート8080でアプリケーションを実行
CMD ["./main"]

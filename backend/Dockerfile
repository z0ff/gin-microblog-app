FROM golang:1.23.0-bookworm
LABEL authors="z0ff"

# パッケージの更新
RUN apt update && apt upgrade -y

# gitのインストール
RUN apt install -y git

# ワーキングディレクトリの設定
WORKDIR /app

# ファイルのコピー
# COPY . .

# 依存パッケージのダウンロード
# RUN go mod download

# ビルド
#RUN go build -o /go/bin/hello-go

# ホットリロード
RUN go install github.com/air-verse/air@latest

# コンテナ起動時のコマンド
#ENTRYPOINT ["/go/bin/hello-go"]
# CMD ["air"]
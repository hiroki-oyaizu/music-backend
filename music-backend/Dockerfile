# 使用するベースイメージ
FROM golang:1.16

# 作業ディレクトリを設定
WORKDIR /go/src/app

# go.modとgo.sumをコピー
COPY go.mod go.sum ./

# ソースコードをコピー
COPY . .

# Goアプリケーションをビルド
RUN go build -o main main.go

# アプリケーションを実行
CMD [ "./main" ]

# 開発環境用 golang イメージ
FROM golang:1.9 AS build_env

WORKDIR /go/src/github.com/muroya2355/osake/go/

RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/lib/pq

ADD ./go .
ENV CGO_ENABLED 0
RUN go build -o osake

# 実行環境用 alipne イメージ
FROM alpine:3.8 AS run_env

RUN mkdir -p /opt/osake/view
ADD  ./go/view /opt/osake/view

COPY --from=build_env /go/src/github.com/muroya2355/osake/go/osake /opt/osake/osake


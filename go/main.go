package main

import (
	"net/http"

	"github.com/muroya2355/denki/controller"
	"github.com/muroya2355/denki/utils"
)

// メイン関数
func main() {

	utils.Init()

	// マルチプレクサの生成
	mux := http.NewServeMux()

	// ハンドラ関数の登録
	mux.HandleFunc("/login", controller.Login)               // ログインページの表示
	mux.HandleFunc("/authenticate", controller.Authenticate) // ログイン認証

	// サーバの生成、マルギプレクサの登録
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

package main

import (
	"html/template"
	"log"
	"net/http"
)

// GET login : ログインページの表示
func login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

// メイン関数
func main() {

	// マルチプレクサの生成
	mux := http.NewServeMux()

	// ルートURL をハンドラ関数にリダイレクト
	//mux.HandleFunc("/", index)

	// ハンドラ関数の登録
	mux.HandleFunc("/login", login)               // ログインページの表示
	mux.HandleFunc("/authenticate", authenticate) // ログイン認証

	// サーバの生成、マルギプレクサの登録
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

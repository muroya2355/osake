package utils

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// CheckCookie : ログインし、クッキーを取得しているかを確認、取得していなかった場合はログイン画面にリダイレクトさせる
func CheckCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {

	// クッキーの取得
	cookie, err := r.Cookie("auth")
	// クッキーを取得していない、またはValue にユーザID が設定されていない場合、ログイン画面にリダイレクト
	if err == http.ErrNoCookie || cookie.Value == "" {
		http.Redirect(w, r, "/login", 302)
	} else if err != nil {
		log.Fatal(err)
	}

	return cookie
}

// Db : データベースに接続するためのハンドラ
var Db *sql.DB

// DbInit : Db の初期化
func DbInit() {

	var err error
	Db, err = sql.Open("postgres", "host=postgres-container port=5432 user=osake_user dbname=osake password=osake_user sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

}

package controller

import (
	"html/template"
	"log"
	"net/http"
)

// POST authenticate : ユーザの認証
func authenticate(w http.ResponseWriter, r *http.Request) {
	// リクエストの解析
	r.ParseForm()

	if true {
		// 認証に成功した場合

		// ユーザ構造体 user の生成
		user := User{r.PostForm["loginid"][0], r.PostForm["password"][0]}

		// ログイン成功画面テンプレートを解析
		tmpl, err := template.ParseFiles("templates/loginsuccessful.html")
		if err != nil {
			log.Fatal(err)
		}
		// ログイン成功画面に user を渡して遷移
		err = tmpl.Execute(w, user)
		if err != nil {
			log.Fatal(err)
		}

	} else {

		// 認証に失敗した場合
		// ログイン画面に遷移
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

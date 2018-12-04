package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/muroya2355/denki/model"
)

// Login : GET ログインページの表示
func Login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/login.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

// Authenticate : POST ユーザの認証
func Authenticate(w http.ResponseWriter, r *http.Request) {
	// リクエストの解析
	r.ParseForm()

	// ログインIDを基にユーザを検索
	user := model.SelectByID(r.PostForm["loginid"][0])

	// 入力されたログインIDのユーザが存在するか、パスワードが一致するか確認
	if user.Loginid != "" && user.Password == r.PostForm["password"][0] {
		// 認証に成功した場合

		// ログイン成功画面テンプレートを解析
		tmpl, err := template.ParseFiles("view/loginsuccessful.html")
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
		tmpl, err := template.ParseFiles("view/login.html")
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

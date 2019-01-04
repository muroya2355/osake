package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/muroya2355/osake/go/model"
)

// Login : GET ログインページの表示
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tmpl, err := template.ParseFiles("view/login.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

// Logout : GET ログアウト、ログインページにリダイレクト
func Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// クッキーの変更
	cookie := http.Cookie{
		Name:     "auth",
		Value:    "",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	// ログインページにリダイレクト
	http.Redirect(w, r, "/login", 302)
}

// Authenticate : POST ユーザの認証
func Authenticate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// リクエストの解析
	r.ParseForm()

	// ログインIDを基にユーザを検索
	user := model.SelectByID(r.PostForm["loginid"][0])

	// 入力されたログインIDのユーザが存在するか、パスワードが一致するか確認
	if user.UserID != "" && user.Password == r.PostForm["password"][0] {
		// 認証に成功した場合

		// クッキーの生成
		cookie := http.Cookie{
			Name:     "auth",
			Value:    user.UserID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		// 商品一覧ページにリダイレクト
		http.Redirect(w, r, "/list", 302)

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

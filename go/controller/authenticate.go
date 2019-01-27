package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muroya2355/osake/go/model"
	"github.com/muroya2355/osake/go/utils"
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
	loginid := r.PostForm["loginid"][0]
	password := r.PostForm["password"][0]

	// エラーメッセージ格納リスト
	Errlist := make([]string, 0)

	// 入力チェック
	flag := true
	if utils.IsEmpty(loginid) {
		flag = false
		Errlist = append(Errlist, utils.Error_required1("ユーザID"))
	}
	if !utils.IsStringLengthWithin(loginid, 1, 10) {
		flag = false
		Errlist = append(Errlist, utils.Error_length("ユーザID", 1, 10))
	}
	if !utils.IsHankakuAlphaNum(loginid) {
		flag = false
		Errlist = append(Errlist, utils.Error_required2("ユーザID", "半角英数字"))
	}

	// 入力エラーがなかった場合
	if flag {
		// ログインIDを基にユーザを検索
		user := model.SelectByID(loginid)

		// 入力されたログインIDのユーザが存在するか、パスワードが一致するか確認
		if user.UserID == loginid && user.Password == password {
			// 認証に成功した場合

			// クッキーの生成
			cookie := http.Cookie{
				Name:     "auth",
				Value:    user.UserID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			// 商品一覧ページにリダイレクト
			http.Redirect(w, r, "/goodslist", 302)

		} else {
			// 認証に失敗した場合
			flag = false
			Errlist = append(Errlist, utils.Error_login())
		}
	}
	// エラー発生時
	if !flag {

		// ログイン画面に遷移
		tmpl, err := template.ParseFiles("view/login.html")
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, Errlist)
		if err != nil {
			log.Fatal(err)
		}
	}
}

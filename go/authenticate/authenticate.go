package controller

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/coreos/go-oidc"
	"github.com/julienschmidt/httprouter"
	"github.com/muroya2355/osake/go/model"
	"github.com/muroya2355/osake/go/utils"
	"golang.org/x/oauth2"
)

var once sync.Once

var provider *oidc.Provider
var oauth2Config *oauth2.Config

// OpenID Connect設定
func getConfig() (*oauth2.Config, *oidc.Provider) {
	once.Do(func() {
		var err error
		// ここにissuer情報を設定
		provider, err = oidc.NewProvider(context.Background(), "http://192.168.99.100:18080/auth/realms/demo")
		if err != nil {
			panic(err)
		}
		oauth2Config = &oauth2.Config{
			// ここにクライアントIDとクライアントシークレットを設定
			ClientID:     "testapp",
			ClientSecret: "40abe9e2-2353-4a13-8485-ca4be8883fc2",
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID},
			RedirectURL:  "http://192.168.99.100:18081/callback",
		}
	})
	return oauth2Config, provider
}

// CheckCookie : ログインし、クッキーを取得しているかを確認、取得していなかった場合はログイン画面にリダイレクトさせる
func CheckCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {

	// クッキーの取得
	cookie, err := r.Cookie("auth")
	// クッキーを取得していない、またはValue にユーザID が設定されていない場合、ログイン画面にリダイレクト
	if err == http.ErrNoCookie || cookie.Value == "" {
		config, _ := getConfig()
		url := config.AuthCodeURL("")
		http.Redirect(w, r, url, http.StatusFound)
	} else if err != nil {
		log.Fatal(err)
	}

	return cookie
}

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

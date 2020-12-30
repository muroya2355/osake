package utils

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/coreos/go-oidc"
	"github.com/julienschmidt/httprouter"
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
			ClientSecret: "a5a745a5-99f0-463b-ae71-9f5c399b3b60",
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID},
			RedirectURL:  "http://www.osaketen.com/callback",
		}
	})
	return oauth2Config, provider
}

// CheckCookie : ログインし、クッキーを取得しているかを確認、取得していなかった場合はログイン画面にリダイレクトさせる
func CheckCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {

	// クッキーの取得
	cookie, err := r.Cookie("auth")
	// クッキーを取得していない、またはValue にユーザID が設定されていない場合、keycloakにリダイレクト
	if err == http.ErrNoCookie || cookie.Value == "" {
		config, _ := getConfig()
		url := config.AuthCodeURL("")
		http.Redirect(w, r, url, http.StatusFound)
	} else if err != nil {
		log.Fatal(err)
	}

	return cookie
}

// Callback : ログイン後のリダイレクト
func Callback(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// 認証に成功した場合

	// クッキーの生成
	cookie := http.Cookie{
		Name:     "auth",
		Value:    "johndoe",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	// 商品一覧ページにリダイレクト
	http.Redirect(w, r, "/goodslist", 302)

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

	// 商品一覧ページにリダイレクト試行
	http.Redirect(w, r, "http://192.168.99.100:18080/auth/realms/demo/protocol/openid-connect/logout?redirect_uri=http%3A%2F%2Fwww.osaketen.com%2F", 302)
}

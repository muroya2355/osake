package utils

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/coreos/go-oidc"
	"github.com/julienschmidt/httprouter"
	"github.com/koron/go-dproxy"
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
		provider, err = oidc.NewProvider(context.Background(), "http://auth.192.168.99.100.xip.io:18080/auth/realms/demo")
		if err != nil {
			panic(err)
		}
		oauth2Config = &oauth2.Config{
			// ここにクライアントIDとクライアントシークレットを設定
			ClientID:     "testapp",
			ClientSecret: "a5a745a5-99f0-463b-ae71-9f5c399b3b60",
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID},
			RedirectURL:  "http://www.192.168.99.100.xip.io/callback",
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
	config, provider := getConfig()
	if err := r.ParseForm(); err != nil {
		http.Error(w, "parse form error", http.StatusInternalServerError)
		return
	}
	accessToken, err := config.Exchange(context.Background(), r.Form.Get("code"))
	if err != nil {
		http.Error(w, "Can't get access token", http.StatusInternalServerError)
		return
	}
	rawIDToken, ok := accessToken.Extra("id_token").(string)
	if !ok {
		http.Error(w, "missing token", http.StatusInternalServerError)
		return
	}
	oidcConfig := &oidc.Config{
		ClientID: "testapp",
	}
	verifier := provider.Verifier(oidcConfig)
	idToken, err := verifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		http.Error(w, "id token verify error", http.StatusInternalServerError)
		return
	}
	// IDトークンのクレームをとりあえずダンプ
	// アプリで必要なものはセッションストレージに入れておくと良いでしょう
	idTokenClaims := map[string]interface{}{}
	if err := idToken.Claims(&idTokenClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if uname, err := dproxy.New(idTokenClaims).M("preferred_username").String(); err != nil {
		log.Fatal(err)
	} else {
		// クッキーの生成
		cookie := http.Cookie{
			Name:  "auth",
			Value: uname,
		}
		http.SetCookie(w, &cookie)

		// 商品一覧ページにリダイレクト
		http.Redirect(w, r, "/goodslist", 302)
	}
}

// Logout : GET ログアウト、ログインページにリダイレクト
func Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// クッキーの変更
	cookie := http.Cookie{
		Name:  "auth",
		Value: "",
	}
	http.SetCookie(w, &cookie)

	// 商品一覧ページにリダイレクト試行
	http.Redirect(w, r, "http://auth.192.168.99.100.xip.io:18080/auth/realms/demo/protocol/openid-connect/logout?redirect_uri=http%3A%2F%2Fwww.192.168.99.100.xip.io%2F", 302)
}

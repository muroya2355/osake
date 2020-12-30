package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muroya2355/osake/go/model"
	"github.com/muroya2355/osake/go/utils"
)

// GoodsList : GETメソッド, Goods 検索画面の表示
func GoodsList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// クッキーの確認
	utils.CheckCookie(w, r)

	// /list へリダイレクトした場合、検索パラメータはない
	query := ""

	// URL から検索パラメータの取得
	q := r.URL.Query()
	// query に値がある場合、値を取得
	if q != nil {
		query = q.Get("query")
	}

	// Goods を検索
	goodsList := model.SearchGoods(query)

	// 商品一覧画面に遷移
	tmpl, err := template.ParseFiles("view/goodslist.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, goodsList)
	if err != nil {
		log.Fatal(err)
	}

}

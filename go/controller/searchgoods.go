package controller

import (
	"html/template"
	"log"
	"net/http"

	"denki/go/model"

	"github.com/julienschmidt/httprouter"
)

// GoodsList : Goods 検索画面の表示
func GoodsList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tmpl, err := template.ParseFiles("view/goodslist.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

// SearchGoods : Goods 検索／検索結果の表示
func SearchGoods(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// リクエストの解析
	r.ParseForm()

	// 検索文字列から Goods を検索
	goodsList := model.SearchGoods(r.PostForm["goodsname"][0])

	tmpl, err := template.ParseFiles("view/goodslist.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, goodsList)
	if err != nil {
		log.Fatal(err)
	}
}

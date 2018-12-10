package controller

import (
	"denki/go/model"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func DeleteGoods(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// フォームの解析
	r.ParseForm()

	// GoodsID の取得
	goodsID, err := strconv.Atoi(r.PostForm["goodsid"][0])
	if err != nil {
		log.Fatal(err)
	}

	// 商品情報の削除
	model.DeleteGoods(goodsID)

	// 商品一覧画面にリダイレクト
	http.Redirect(w, r, "/default", 301)
}

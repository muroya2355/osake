package controller

import (
	"denki/go/model"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GoodsDetail : 商品詳細情報の表示
func GoodsDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// 商品詳細用の構造体を定義
	type DisplayGoodsDetail struct {
		Goods   model.Goods
		Classes []model.Class
		Makers  []model.Maker
	}

	// DisplayGoodsDetail 構造体を宣言
	var displayGoodsDetail DisplayGoodsDetail
	// クラス・メーカーの取得、格納
	displayGoodsDetail.Classes = model.SelectAllClass()
	displayGoodsDetail.Makers = model.SelectAllMaker()

	// フォームの解析
	r.ParseForm()

	// GoodsIDの取得
	goodsid, err := strconv.Atoi(r.PostForm["goodsid"][0])
	if err != nil {
		log.Fatal(err)
	}

	// DBに接続、商品情報の取得
	displayGoodsDetail.Goods = model.GetGoods(goodsid)

	// 商品詳細画面の表示
	tmpl, err := template.ParseFiles("view/goodsdetail.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, displayGoodsDetail)
	if err != nil {
		log.Fatal(err)
	}
}

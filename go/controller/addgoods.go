package controller

import (
	"denki/go/model"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// AddGoods : 商品の追加
func AddGoods(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// フォームの解析
	r.ParseForm()

	// 追加用の Goods の生成、パラメータの代入
	var goods model.Goods
	goods.GoodsName = r.PostForm["goodsname"][0]
	goods.ClassID, _ = strconv.Atoi(r.PostForm["classid"][0])
	goods.MakerID, _ = strconv.Atoi(r.PostForm["makerid"][0])
	goods.ModelNumber = r.PostForm["modelnumber"][0]
	goods.Specs = r.PostForm["specs"][0]
	goods.IndicatedPrice, _ = strconv.Atoi(r.PostForm["indicatedprice"][0])
	goods.PurchasePrice, _ = strconv.ParseFloat(r.PostForm["purchaseprice"][0], 64)
	goods.Stock, _ = strconv.Atoi(r.PostForm["stock"][0])
	goods.Deleted, _ = strconv.ParseBool(r.PostForm["isdeleted"][0])

	// DBに接続、商品追加
	goods.AddGoods()

	// 商品一覧画面にリダイレクト
	http.Redirect(w, r, "/default", 301)
}

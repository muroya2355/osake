package controller

import (
	"net/http"
	"strconv"

	"github.com/muroya2355/osake/go/model"
	"github.com/muroya2355/osake/go/utils"

	"github.com/julienschmidt/httprouter"
)

// UpdateGoods : 商品の追加
func UpdateGoods(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// クッキーの確認
	utils.CheckCookie(w, r)

	// フォームの解析
	r.ParseForm()

	// 追加用の Goods の生成、パラメータの代入
	var goods model.Goods
	goods.GoodsID, _ = strconv.Atoi(r.PostForm["goodsid"][0])
	goods.GoodsName = r.PostForm["goodsname"][0]
	goods.ClassID, _ = strconv.Atoi(r.PostForm["classid"][0])
	goods.MakerID, _ = strconv.Atoi(r.PostForm["makerid"][0])
	goods.IndicatedPrice, _ = strconv.Atoi(r.PostForm["indicatedprice"][0])
	goods.PurchasePrice, _ = strconv.ParseFloat(r.PostForm["purchaseprice"][0], 64)
	goods.Stock, _ = strconv.Atoi(r.PostForm["stock"][0])

	// DBに接続、商品追加
	goods.UpdateGoods()

	// 商品一覧画面にリダイレクト
	http.Redirect(w, r, "/list", 301)
}

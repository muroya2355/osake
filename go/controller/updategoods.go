package controller

import (
	"log"
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
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// 入力値チェック
	// エラーメッセージ格納リスト
	Errlist := make([]string, 0)

	// 入力チェック
	flag := true
	if utils.IsEmpty(r.PostForm["goodsname"][0]) {
		flag = false
		Errlist = append(Errlist, utils.Error_required1("ユーザID"))
	}
	if utils.IsEmpty(r.PostForm["indicatedprice"][0]) {
		flag = false
		Errlist = append(Errlist, utils.Error_required1("ユーザID"))
	}
	if utils.IsEmpty(r.PostForm["purchaseprice"][0]) {
		flag = false
		Errlist = append(Errlist, utils.Error_required1("ユーザID"))
	}
	if utils.IsEmpty(r.PostForm["stock"][0]) {
		flag = false
		Errlist = append(Errlist, utils.Error_required1("ユーザID"))
	}

	// 追加用の Goods の生成、パラメータの代入
	var goods model.Goods
	goods.GoodsID, _ = strconv.Atoi(r.PostForm["goodsid"][0])
	goods.GoodsName = r.PostForm["goodsname"][0]
	goods.ClassID, _ = strconv.Atoi(r.PostForm["classid"][0])
	goods.MakerID, _ = strconv.Atoi(r.PostForm["makerid"][0])
	goods.IndicatedPrice, err = strconv.Atoi(r.PostForm["indicatedprice"][0])
	if err != nil {
		flag = false
	}
	goods.PurchasePrice, err = strconv.ParseFloat(r.PostForm["purchaseprice"][0], 64)
	if err != nil {
		flag = false
	}
	goods.Stock, err = strconv.Atoi(r.PostForm["stock"][0])
	if err != nil {
		flag = false
	}

	if flag {
		// DBに接続、商品追加
		goods.UpdateGoods()
	} else {
		http.Redirect(w, r, "/list", 301)
	}

	// 商品一覧画面にリダイレクト
	http.Redirect(w, r, "/list", 301)
}

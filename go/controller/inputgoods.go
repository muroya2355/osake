package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/muroya2355/osake/go/model"
	"github.com/muroya2355/osake/go/utils"

	"github.com/julienschmidt/httprouter"
)

// InputGoods : Class と Goods の全検索、商品追加画面の表示
func InputGoods(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// クッキーの確認
	utils.CheckCookie(w, r)

	// 全クラス、全メーカを格納する構造体
	type Displayinputgoods struct {
		Classes []model.Class
		Makers  []model.Maker
	}

	// Displayinputgoods 構造体の生成
	var displayinputgoods Displayinputgoods
	// クラス・メーカーの検索、格納
	displayinputgoods.Classes = model.SelectAllClass()
	displayinputgoods.Makers = model.SelectAllMaker()

	// 商品追加画面の表示
	tmpl, err := template.ParseFiles("view/inputgoods.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, displayinputgoods)
	if err != nil {
		log.Fatal(err)
	}
}

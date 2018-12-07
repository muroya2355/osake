package model

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muroya2355/denki/go/utils"
)

type Goods struct {
	GoodsId        int
	GoodsName      string
	ClassName      string
	MakerName      string
	ModelNumber    string
	Specs          string
	IndicatedPrice int
	PurchasePrice  float64
	Stock          int
	Deleted        bool
}

// func SearchGoods(query string) []Goods { }

// GetGoods : /login/:id  GoodsId から Goods を取得し、JSON形式で返す
func GetGoods(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// URI から id の取得
	goodsid := p.ByName("id")

	// SQL 文の構築
	sql := "SELECT GOODS.goods_id, GOODS.goods_name, CLASS.class_name, " +
		"MAKER.maker_name, GOODS.model_number, GOODS.specs, GOODS.indicated_price " +
		"GOODS.purchase_price, GOODS.stock, GOODS.is_deleted FROM GOODS " +
		"INNER JOIN CLASS ON WHERE CLASS.class_id = GOODS.class_id " +
		"INNER JOIN MAKER ON WHERE MAKER.maker_id = GOODS.maker_id " +
		"WHERE GOODS_id = $1;"

	// preparedstatement の生成
	pstatement, err := utils.Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// 結果を格納する Goods 構造体
	var goods Goods

	// パラメータに id を埋め込み SQL 文の実行、結果を goods 構造体に格納する
	err = pstatement.QueryRow(goodsid).Scan(&goods.GoodsId, &goods.GoodsName, &goods.ClassName, &goods.MakerName, &goods.ModelNumber, &goods.Specs, &goods.IndicatedPrice, &goods.PurchasePrice, &goods.Stock, &goods.Deleted)
	if err != nil {
		log.Fatal(err)
	}

}

// func GetGoodsList() []Goods { }
func (goods *Goods) AddGoods() {

}
func (goods *Goods) UpdateGoods() {

}

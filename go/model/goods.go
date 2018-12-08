package model

import (
	"log"

	"denki/go/utils"
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

func GetGoods(id int) Goods {

	// SQL 文の構築
	sql := "SELECT GOODS.goods_id, GOODS.goods_name, CLASS.class_name, " +
		"MAKER.maker_name, GOODS.model_number, GOODS.specs, GOODS.indicated_price " +
		"GOODS.purchase_price, GOODS.stock, GOODS.is_deleted FROM GOODS " +
		"INNER JOIN CLASS ON WHERE CLASS.class_id = GOODS.class_id " +
		"INNER JOIN MAKER ON WHERE MAKER.maker_id = GOODS.maker_id " +
		"WHERE GOODS.goods_id = $1;"

	// preparedstatement の生成
	pstatement, err := utils.Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// 結果を格納する Goods 構造体
	var goods Goods

	// パラメータに id を埋め込み SQL 文の実行、結果を goods 構造体に格納する
	err = pstatement.QueryRow(id).Scan(&goods.GoodsId, &goods.GoodsName, &goods.ClassName, &goods.MakerName, &goods.ModelNumber, &goods.Specs, &goods.IndicatedPrice, &goods.PurchasePrice, &goods.Stock, &goods.Deleted)
	if err != nil {
		log.Fatal(err)
	}

	return goods
}

// SearchGoods : goodsname より Goods を検索
func SearchGoods(goodsname string) []Goods {

	// SQL 文の構築
	sql := "SELECT GOODS.goods_id, GOODS.goods_name, CLASS.class_name, MAKER.maker_name, GOODS.model_number, GOODS.specs, GOODS.indicated_price, GOODS.purchase_price, GOODS.stock, GOODS.is_deleted FROM GOODS JOIN CLASS ON CLASS.class_id = GOODS.class_id JOIN MAKER ON MAKER.maker_id = GOODS.maker_id WHERE GOODS.goods_name LIKE $1;"

	// preparedstatement の生成
	pstatement, err := utils.Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// 結果を格納する Goods 構造体
	var goodsList []Goods

	// パラメータに id を埋め込み SQL 文の実行、結果を Rows ポインタに格納する
	rows, err1 := pstatement.Query("%" + goodsname + "%")
	if err1 != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var goods Goods
		err = rows.Scan(&goods.GoodsId, &goods.GoodsName, &goods.ClassName, &goods.MakerName, &goods.ModelNumber, &goods.Specs, &goods.IndicatedPrice, &goods.PurchasePrice, &goods.Stock, &goods.Deleted)
		if err != nil {
			log.Fatal(err)
		}
		goodsList = append(goodsList, goods)
	}

	return goodsList

}

// func GetGoodsList() []Goods { }
func (goods *Goods) AddGoods() {

}
func (goods *Goods) UpdateGoods() {

}

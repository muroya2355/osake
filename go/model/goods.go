package model

import (
	"log"

	"denki/go/utils"
)

// Goods : 商品情報
type Goods struct {
	GoodsID        int
	GoodsName      string
	ClassName      string
	MakerName      string
	ModelNumber    string
	Specs          string
	IndicatedPrice int
	PurchasePrice  float64
	Stock          int
	Deleted        bool
	ClassID        int
	MakerID        int
}

// GetGoods : id より商品情報を取得
func GetGoods(id int) Goods {

	// SQL 文の構築
	sql := "SELECT GOODS.goods_id, GOODS.goods_name, GOODS.class_id, GOODS.maker_id, CLASS.class_name,MAKER.maker_name, GOODS.model_number, GOODS.specs, GOODS.indicated_price, GOODS.purchase_price, GOODS.stock, GOODS.is_deleted FROM GOODS " +
		"INNER JOIN CLASS ON CLASS.class_id = GOODS.class_id " +
		"INNER JOIN MAKER ON MAKER.maker_id = GOODS.maker_id " +
		"WHERE GOODS.goods_id = $1;"

	// preparedstatement の生成
	pstatement, err := utils.Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// 結果を格納する Goods 構造体
	var goods Goods

	// パラメータに id を埋め込み SQL 文の実行、結果を goods 構造体に格納する
	err = pstatement.QueryRow(id).Scan(&goods.GoodsID, &goods.GoodsName, &goods.ClassID, &goods.MakerID, &goods.ClassName, &goods.MakerName, &goods.ModelNumber, &goods.Specs, &goods.IndicatedPrice, &goods.PurchasePrice, &goods.Stock, &goods.Deleted)
	if err != nil {
		log.Fatal(err)
	}

	return goods
}

// SearchGoods : goodsname より Goods を検索
func SearchGoods(goodsname string) []Goods {

	// SQL 文の構築
	sql := "SELECT GOODS.goods_id, GOODS.goods_name, CLASS.class_name, MAKER.maker_name, GOODS.model_number, GOODS.specs, GOODS.indicated_price, GOODS.purchase_price, GOODS.stock, GOODS.is_deleted FROM GOODS JOIN CLASS ON CLASS.class_id = GOODS.class_id JOIN MAKER ON MAKER.maker_id = GOODS.maker_id WHERE GOODS.goods_name LIKE $1 ORDER BY GOODS.goods_id ASC;"

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
		err = rows.Scan(&goods.GoodsID, &goods.GoodsName, &goods.ClassName, &goods.MakerName, &goods.ModelNumber, &goods.Specs, &goods.IndicatedPrice, &goods.PurchasePrice, &goods.Stock, &goods.Deleted)
		if err != nil {
			log.Fatal(err)
		}
		goodsList = append(goodsList, goods)
	}

	return goodsList

}

// AddGoods : 商品を追加、GoodsID を取得する
func (goods *Goods) AddGoods() {

	// SQL 文の構築
	sql := "INSERT INTO GOODS (" +
		"goods_id, goods_name, class_id, maker_id, model_number, specs, indicated_price, purchase_price, stock, is_deleted, update_super_visor_id, update_date, update_version_id) " +
		"SELECT  MAX(goods_id)+1, $1, $2, $3, $4, $5, $6, $7, $8, $9, 'a', current_timestamp, 1 FROM GOODS RETURNING goods_id;"

	// preparedstatement の生成
	pstatement, err := utils.Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer pstatement.Close()

	// SQL文にパラメータを埋め込み実行、GoodsIDを取得
	err = pstatement.QueryRow(goods.GoodsName, goods.ClassID, goods.MakerID, goods.ModelNumber, goods.Specs, goods.IndicatedPrice, goods.PurchasePrice, goods.Stock, goods.Deleted).Scan(&goods.GoodsID)
	if err != nil {
		log.Fatal(err)
	}

}

// UpdateGoods : 商品情報の更新
func (goods *Goods) UpdateGoods() {

	// SQL文の構築
	sql := "UPDATE GOODS " +
		"SET goods_name=$1, class_id=$2, maker_id=$3, model_number=$4, specs=$5, indicated_price=$6, purchase_price=$7, stock=$8, is_deleted=$9, update_super_visor_id='a', update_date=current_timestamp, update_version_id=1 " +
		"WHERE goods_id=$10;"

	// preparedstatement の生成
	pstatement, err := utils.Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// SQL文の実行
	_, err = pstatement.Exec(goods.GoodsName, goods.ClassID, goods.MakerID, goods.ModelNumber, goods.Specs, goods.IndicatedPrice, goods.PurchasePrice, goods.Stock, goods.Deleted, goods.GoodsID)
	if err != nil {
		log.Fatal(err)
	}

}

// DeleteGoods : 商品情報の削除
func DeleteGoods(goodsid int) {

	// SQL文の構築
	sql := "DELETE FROM GOODS WHERE goods_id = $1"

	// preparedstatement の生成
	pstatement, err := utils.Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// SQL文にパラメータを埋め込み、SQL文の実行
	_, err = pstatement.Exec(goodsid)
	if err != nil {
		log.Fatal(err)
	}

}

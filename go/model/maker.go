package model

import (
	"denki/go/utils"
	"log"
)

type Maker struct {
	MakerID   int
	MakerName string
}

func SelectAllMaker() []Maker {

	// SQL文の構築
	sql := "SELECT maker_id, maker_name FROM MAKER;"

	// preparedstatement の生成
	pstatement, err := utils.Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// 検索結果格納用の Maker スライスを宣言
	var makers []Maker

	// SQL文の実行、結果を Maker スライスに格納
	rows, err1 := pstatement.Query()
	if err1 != nil {
		log.Fatal(err1)
	}

	for rows.Next() {
		var maker Maker
		rows.Scan(&maker.MakerID, &maker.MakerName)
		makers = append(makers, maker)
	}

	// Maker スライスの返却
	return makers
}

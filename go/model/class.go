package model

import (
	"log"

	"github.com/muroya2355/osake/go/utils"
)

// Class : 分類
type Class struct {
	ClassID   int
	ClassName string
}

func SelectAllClass() []Class {

	// SQL文の構築
	sql := "SELECT class_id, class_name FROM CLASS;"

	// preparedstatement の生成
	pstatement, err := utils.Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// 検索結果格納用の Class スライスを宣言
	var classes []Class

	// SQL文の実行、結果を Class スライスに格納
	rows, err := pstatement.Query()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var class Class
		err = rows.Scan(&class.ClassID, &class.ClassName)
		if err != nil {
			log.Fatal(err)
		}
		classes = append(classes, class)
	}

	// Class スライスの返却
	return classes
}

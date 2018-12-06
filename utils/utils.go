package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Db : データベースに接続するためのハンドラ
var Db *sql.DB

// Init : Db の初期化
func Init() {

	var err error
	Db, err = sql.Open("postgres", "user=denki_user dbname=denki password=denki_user sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

}

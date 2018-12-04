package utils

import(
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

// データベースに接続するためのハンドラ
var Db *sql.DB

// Db の初期化
func init() {

	var err error
	Db, err = sql.Open("postgres", "user=denki_user, dbname=denki, password=denki_user sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

}
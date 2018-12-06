package model

import (
	"log"

	"github.com/muroya2355/denki/utils"
)

// User : ログインユーザ
type User struct {
	Userid   string
	Password string
}

// SelectByID : ログインIDを基にユーザを検索する
func SelectByID(loginid string) User {

	// SQL 文の構築
	sql := "SELECT super_visor_id, super_visor_password FROM SUPER_VISOR WHERE super_visor_id = $1;"

	// PreparedStatement の作成
	pstatement, err := utils.Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// 検索結果格納先の User 構造体を作成
	var user User

	// pstatement のパラメータに loginid を埋め込み実行、検索結果を user のフィールドに埋め込む
	err = pstatement.QueryRow(loginid).Scan(&user.Userid, &user.Password)
	if err != nil {
		log.Fatal(err)
	}

	// 検索した user を返却
	return user
}

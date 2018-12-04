package model

// User : ログインユーザ
type User struct {
	Loginid  string
	Password string
}

// SelectByID : ログインIDを基にユーザを検索する
func SelectByID(login string) User {

	user := User{
		Loginid:  "a",
		Password: "password1",
	}
	return user
}

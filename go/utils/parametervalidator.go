package utils

import "regexp"
import "strconv"

// IsEmpty : 引数が空文字のとき、true を返す
func IsEmpty(parameter string) bool {
	return parameter == ""
}

// IsStringLengthWithin : 引数の文字列長がmin~maxである場合、true を返す
func IsStringLengthWithin(parameter string, min int, max int) bool {
	if parameter == "" {
		return true
	}
	if max < min {
		return false
	}
	return min <= len(parameter) && len(parameter) <= max
}

// IsHankakuAlphaNum : 引数の文字列が半角英数字のみの場合、trueを返す
func IsHankakuAlphaNum(parameter string) bool {
	if parameter == "" {
		return true
	}
	flag := true
	for _, char := range parameter {
		f, _ := regexp.MatchString("[A-Za-z0-9]", string(char))
		flag = flag && f
	}
	return flag
}

// エラーメッセージ一覧
func ErrorLogin() string {
	return "ログインに失敗しました。ユーザID、パスワードを確認してください。"
}
func ErrorRequired1(a string) string {
	return a + "を入力してください。"
}
func ErrorRequired2(a, b string) string {
	return a + "には" + b + "を入力してください。"
}
func ErrorLength(a string, from, to int) string {
	return a + "には " + strconv.Itoa(from) + " 文字以上、" + strconv.Itoa(to) + " 文字以下で入力してください。"
}
func ErrorSelect(a string) string {
	return a + "を選択してください。"
}
func ErrorLock() string {
	return "該当の商品は他のユーザに更新されました。再度、商品詳細画面を表示してから、更新業務を行ってください。"
}
func ErrorSystem() string {
	return "システムエラーが発生しました。"
}

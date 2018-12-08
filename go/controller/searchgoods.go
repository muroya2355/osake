package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GoodsList : GET ログインページの表示
func GoodsList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tmpl, err := template.ParseFiles("view/login.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

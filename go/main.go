package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/muroya2355/osake/go/controller"
	"github.com/muroya2355/osake/go/utils"
)

// Index : ログインページにリダイレクト
func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// ログインページにリダイレクト
	http.Redirect(w, r, "/login", 302)
}

// メイン関数
func main() {

	// DB 接続用ハンドラの宣言・初期化
	utils.DbInit()

	// HTTPルーターを初期化
	router := httprouter.New()

	// ハンドラ関数の登録
	router.GET("/", Index)                                // ログインページにリダイレクト
	router.GET("/login", controller.Login)                // ログインページの表示
	router.GET("/logout", controller.Logout)              // ログアウト
	router.POST("/authenticate", controller.Authenticate) // ログイン認証
	router.GET("/goodslist", controller.GoodsList)        // 商品リスト画面の表示
	router.GET("/inputgoods", controller.InputGoods)      // 商品追加画面の表示
	router.POST("/addgoods", controller.AddGoods)         // 商品の追加
	router.GET("/goods/:id", controller.GoodsDetail)      // 商品詳細画面の表示
	router.POST("/updategoods", controller.UpdateGoods)   // 商品情報の更新
	router.POST("/deletegoods", controller.DeleteGoods)   // 商品情報の削除

	// サーバの生成、マルチプレクサの登録
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

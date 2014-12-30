package urls

import (
	"github.com/gorilla/mux"
	"net/http"
	cntl "../controllers"
)

func Run() {
	router := mux.NewRouter()
	// テンプレートのため
	/*
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		ren.HTML(w, http.StatusOK, "index", nil)
	})
	*/

	router.HandleFunc("/", cntl.RootFunc)
	router.HandleFunc("/products/{key}", cntl.ProductsHandler)
	router.HandleFunc("/article/{category}", cntl.ArticlesCategoryHandler)
	// 正規表現にマッチしない場合は404
	// 最後の/がないと直前のパラメータはなかったことになる
	router.HandleFunc("/article/{category}/{id:[0-9]+}/", cntl.ArticleHandler).Methods("GET")

	// 404のときのハンドラ
	// このようにすれば独自に404用のハンドラを定義できるようです。
	router.NotFoundHandler = http.HandlerFunc(cntl.NotFoundHandler)

	// 後に必須
	http.Handle("/", router)
	// port番号
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products/{key}", ProductsHandler)
	r.HandleFunc("/article/{category}", ArticlesCategoryHandler)
	// 正規表現にマッチしない場合は404
	// 最後の/がないと直前のパラメータはなかったことになる
	r.HandleFunc("/article/{category}/{id:[0-9]+}/", ArticleHandler).Methods("GET")

	// 404のときのハンドラ
	// このようにすれば独自に404用のハンドラを定義できるようです。
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for k, v := range vars {
		fmt.Println(k, v)
	}

	w.Write([]byte("ProductsHandler"))
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for k, v := range vars {
		fmt.Println(k, v)
	}

	w.Write([]byte("ArticlesCategoryHandler"))
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for k, v := range vars {
		fmt.Println(k, v)
	}

	w.Write([]byte("ArticleHandleer"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

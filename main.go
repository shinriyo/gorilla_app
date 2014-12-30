package main

import (
	"github.com/gorilla/mux"
	"net/http"
	urls "./urls"
	cntl "./controllers"
)

func main() {
	router := mux.NewRouter()
	urls.Run()
	router.HandleFunc("/products/{key}", cntl.ProductsHandler)

	// 404のときのハンドラ
	// このようにすれば独自に404用のハンドラを定義できるようです。
	router.NotFoundHandler = http.HandlerFunc(cntl.NotFoundHandler)

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}


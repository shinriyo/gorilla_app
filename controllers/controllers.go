package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

func RootFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MyFunc() called")
	ren := render.New(render.Options{
		// a lot of app specific setup
	})

	ren.HTML(w, http.StatusOK, "index", nil)
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


func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for k, v := range vars {
		fmt.Println(k, v)
	}

	w.Write([]byte("ProductsHandler"))
}


func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

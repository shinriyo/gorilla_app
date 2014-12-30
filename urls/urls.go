package urls

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	cntl "../controllers"
)

func Run() {
	router := mux.NewRouter()
	ren := render.New(render.Options{
		// a lot of app specific setup
	})

	// テンプレートのため
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			ren.HTML(w, http.StatusOK, "index", nil)
		})

	router.HandleFunc("/products/{key}", cntl.ProductsHandler)
	router.HandleFunc("/article/{category}", cntl.ArticlesCategoryHandler)
	// 正規表現にマッチしない場合は404
	// 最後の/がないと直前のパラメータはなかったことになる
	router.HandleFunc("/article/{category}/{id:[0-9]+}/", cntl.ArticleHandler).Methods("GET")
}

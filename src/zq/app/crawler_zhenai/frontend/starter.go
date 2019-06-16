package main

import (
	"net/http"
	"zq/app/crawler_zhenai/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("src/zq/app/crawler_zhenai/frontend/view/")))

	http.Handle("/search",
		controller.CreateSearchResultHandler("src/zq/app/crawler_zhenai/frontend/view/template.html"))
	err := http.ListenAndServe(":9527", nil)
	if err != nil {
		panic(err)
	}
}

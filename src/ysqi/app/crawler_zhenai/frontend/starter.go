package main

import (
	"net/http"
	"ysqi/app/crawler_zhenai/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("src/ysqi/app/crawler_zhenai/frontend/view/")))

	http.Handle("/search",
		controller.CreateSearchResultHandler("src/ysqi/app/crawler_zhenai/frontend/view/template.html"))
	err := http.ListenAndServe(":9527", nil)
	if err != nil {
		panic(err)
	}
}

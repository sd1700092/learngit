package main

import (
	"net/http"

	"imooc.com/learngo/crawler/frontend/controller"
)

func main() {
	//http.Handle("/search", controller.SearchResultHandler{})
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

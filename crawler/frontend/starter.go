package main

import (
	"net/http"

	"imooc.com/learngo/crawler/frontend/controller"
)

func main() {
	http.Handle("/search", controller.SearchResultHandler{})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

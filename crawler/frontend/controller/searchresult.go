package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/olivere/elastic.v5"
	"imooc.com/learngo/crawler/frontend/view"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	fmt.Fprintf(w, "q=%s, from=%d", q, from)
}

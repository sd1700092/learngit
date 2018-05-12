package view

import (
	"os"
	"testing"

	"imooc.com/learngo/crawler/engine"
	"imooc.com/learngo/crawler/frontend/model"
	common "imooc.com/learngo/crawler/model"
)

//func TestTemplate(t *testing.T) {
func TestSearchResultView_Render(t *testing.T) {
	//template := template.Must(template.ParseFiles("template.html"))
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1314495053",
		Type: "zhenai",
		Id:   "108906739",
		Payload: common.Profile{
			Age:        41,
			Height:     158,
			Weight:     48,
			Income:     "3001-5000元",
			Xingzuo:    "处女座",
			Occupation: "公务员",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "四川阿坝",
			Education:  "中专",
			Car:        "未购车",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	//err = template.Execute( /*os.Stdout*/ out, page)
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}

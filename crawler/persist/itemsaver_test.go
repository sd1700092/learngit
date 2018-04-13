package persist

import (
	"context"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"
	"imooc.com/learngo/crawler/model"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:       "风中的蒲公英",
		Gender:     "女",
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
	}
	id, err := save(expected)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.99.100:9200"),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s\n", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	if actual != expected {
		t.Errorf("got %v, expected %v\n", actual, expected)
	}
}

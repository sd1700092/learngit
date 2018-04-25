package parser

import (
	"io/ioutil"
	"testing"

	"imooc.com/learngo/crawler/engine"
	"imooc.com/learngo/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "http://album.zhenai.com/u/1314495053", "风中的蒲公英")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	actual := result.Items[0]
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1314495053",
		Type: "zhenai",
		Id:   "1314495053",
		Payload: model.Profile{
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
		},
	}

	if actual != expected {
		t.Errorf("expected: %+v, but was %+v", expected, actual)
	}
}

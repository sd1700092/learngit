package main

import (
	"context"
	"fmt"
	"reflect"

	"gopkg.in/olivere/elastic.v5"
	"imooc.com/learngo/crawler/engine"
)

func main() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.99.100:9200"),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Search("dating_profile").Query(elastic.NewQueryStringQuery("男 已购房")). /*.From(20)*/ Do(context.Background())
	if err != nil {
		panic(err)
	}
	items := resp.Each(reflect.TypeOf(engine.Item{}))
	for _, item := range items {
		fmt.Println(item)
	}
}

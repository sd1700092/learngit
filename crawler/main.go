package main

import (
	//"net/http"
	//"io/ioutil"
	//"fmt"
	//"golang.org/x/text/transform"
	//"io"
	//"golang.org/x/text/encoding"
	//"bufio"
	//"golang.org/x/net/html/charset"
	//"regexp"
	"imooc.com/learngo/crawler/engine"
	"imooc.com/learngo/crawler/zhenai/parser"
	"imooc.com/learngo/crawler/scheduler"
	"imooc.com/learngo/crawler/persist"
)

func main() {
	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan: persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}

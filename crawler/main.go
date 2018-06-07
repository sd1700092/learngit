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
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

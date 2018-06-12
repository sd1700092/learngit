package parser

import (
	"regexp"

	"imooc.com/learngo/crawler/engine"
)

// <a target="_blank" href="http://city.zhenai.com/beijing">北京</a>
const cityListRe = `<a .* href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		// 不用生成city名字了，如果要的话还可以打log
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: ParseCity})
	}
	return result
}

package parser

import (
	"regexp"

	"imooc.com/learngo/crawler/engine"
)

//<a href="http://album.zhenai.com/u/108415017" target="_blank">惠儿</a>
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	//fmt.Printf("%s\n", contents)
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: func(c []byte) engine.ParseResult {
			return ParseProfile(c, name)
		}})
	}
	return result
}

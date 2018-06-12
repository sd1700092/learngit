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
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: ParseCity})
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		//if i == 10 {
		//	fmt.Printf("only testing %d pages, then break", i)
		//	break
		//}
	}
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}

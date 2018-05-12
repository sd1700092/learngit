package engine

import (
	"log"

	"imooc.com/learngo/crawler/fetcher"
)

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %s", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body, r.Url), nil
}

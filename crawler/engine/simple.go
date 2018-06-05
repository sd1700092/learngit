package engine

import (
	"imooc.com/learngo/crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := e.worker(r)
		if err!=nil{
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func (SimpleEngine) worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %s", r.Url, err)
		return ParseResult{}, nil
	}
	return r.ParserFunc(body), nil
}

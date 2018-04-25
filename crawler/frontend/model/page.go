package model

import "imooc.com/learngo/crawler/engine"

type SearchResult struct {
	Hits     int
	Start    int
	Items    []engine.Item
	Query    int
	PrevFrom int
	NextFrom int
}

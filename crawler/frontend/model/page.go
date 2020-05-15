package model

import "chuanshan.github.com/learngo4p/crawler/engine"

type SearchResult struct {
	Hints int
	Start int
	Items []engine.Item
}

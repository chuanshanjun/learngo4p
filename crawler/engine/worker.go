package engine

import (
	"log"

	"chuanshan.github.com/learngo4p/crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetch ulr %v", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	//return r.ParserFunc(body, r.Url), nil
	return r.Parser.Parse(body, r.Url), nil
}

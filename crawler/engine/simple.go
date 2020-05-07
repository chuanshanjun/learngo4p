package engine

import (
	"fmt"
	"log"

	"chuanshan.github.com/learngo4p/crawler/fetcher"
)

type SimpleEngine struct {
}

func (s *SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	// 先把所有的request装载起来
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			log.Printf("Fecth err")
			continue
		}
		requests = append(requests, parseResult.Requests...)
		// 打印item
		for _, item := range parseResult.Items {
			fmt.Printf("Got item %v\n", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetch ulr %v", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}

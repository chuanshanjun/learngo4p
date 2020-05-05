package engine

import (
	"fmt"
	"log"

	"chuanshan.github.com/learngo4p/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// 拿出1个request
		r := requests[0]
		// 从切片中减少一个
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			fmt.Printf("got item :%s\n", item)
		}
	}
}

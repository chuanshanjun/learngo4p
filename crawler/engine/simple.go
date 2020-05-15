package engine

import (
	"fmt"
	"log"
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

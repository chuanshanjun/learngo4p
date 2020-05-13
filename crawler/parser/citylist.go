package parser

import (
	"fmt"
	"regexp"

	"chuanshan.github.com/learngo4p/crawler/engine"
)

const cityListRe = `<a href="(http://m.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	cityCounts := 0
	for _, m := range matches {
		//result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		cityCounts++
		if cityCounts == 10 {
			break
		}
	}

	//fmt.Printf("Matched found: %d\n", len(matches))
	return result
}

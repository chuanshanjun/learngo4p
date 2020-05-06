package parser

import (
	"regexp"

	"chuanshan.github.com/learngo4p/crawler/engine"
)

const cityRe = `<a href="(http://m.zhenai.com/u/[0-9]+)".*\s+(.+)\s+<span`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	macthes := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range macthes {
		name := string(m[2])
		result.Items = append(result.Items, "User: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
		//fmt.Printf("User: %s, URL: %s\n", m[2], m[1])
	}
	return result
}

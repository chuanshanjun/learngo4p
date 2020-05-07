package parser

import (
	"regexp"

	"chuanshan.github.com/learngo4p/crawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://m.zhenai.com/u/[0-9]+)".*\s+(.+)\s+<span`)
	// 原先只是用上海匹配的
	//cityUrlRe = regexp.MustCompile(`href="(http://m.zhenai.com/zhenghun/shanghai/[^"]+)"`)
	cityUrlRe = regexp.MustCompile(`href="(http://m.zhenai.com/zhenghun/[^"]+/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
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

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}

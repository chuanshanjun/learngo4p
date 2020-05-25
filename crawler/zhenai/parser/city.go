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

// 咱url不用也可以_
func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//url := string(m[1])
		// 我们此时已经不是一个闭包，而是一个函数调用，那就不需要把name
		// 给拷贝出来，函数调用的时候参数本身就是一个拷贝
		//name := string(m[2])
		//result.Items = append(result.Items, "User: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParserFunc: func(c []byte) engine.ParseResult {
			//	return ParseProfile(c, url, name)
			//},
			Parser: NewProfileParser(string(m[2])),
		})
		//fmt.Printf("User: %s, URL: %s\n", m[2], m[1])
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}

package parser

import (
	"regexp"

	"chuanshan.github.com/learngo4p/crawler/engine"
)

const cityListRe = `<a href="(http://m.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	cityCounts := 0
	for _, m := range matches {
		//result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// 函数parseCity写完后，我们用函数NewFuncParser包装下
			// 反序列化，最终我们最终在网络上传输的是不是ParseCity函数，而是ParseCity字符串
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		cityCounts++
		if cityCounts == 10 {
			break
		}
	}

	//fmt.Printf("Matched found: %d\n", len(matches))
	return result
}

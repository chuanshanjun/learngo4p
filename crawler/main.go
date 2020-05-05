package main

import (
	"chuanshan.github.com/learngo4p/crawler/engine"
	"chuanshan.github.com/learngo4p/crawler/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://m.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

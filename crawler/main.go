package main

import (
	"chuanshan.github.com/learngo4p/crawler/engine"
	"chuanshan.github.com/learngo4p/crawler/parser"
	"chuanshan.github.com/learngo4p/crawler/scheduler"
)

func main() {
	e := engine.Concurrent{
		// 用指针...
		WorkerCount: 10,
		Scheduler:   &scheduler.QueuedScheduler{},
	}
	e.Run(engine.Request{
		Url:        "http://m.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//e := engine.SimpleEngine{}
	//e.Run(engine.Request{
	//	Url:        "http://m.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	//engine.Run(engine.Request{
	//	Url:        "http://m.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
}

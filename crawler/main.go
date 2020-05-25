package main

import (
	"chuanshan.github.com/learngo4p/crawler/engine"
	"chuanshan.github.com/learngo4p/crawler/persist"
	"chuanshan.github.com/learngo4p/crawler/scheduler"
	"chuanshan.github.com/learngo4p/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.Concurrent{
		// 用指针...
		WorkerCount: 100,
		Scheduler:   &scheduler.QueuedScheduler{},
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url: "http://m.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList, "ParseCityList"),
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

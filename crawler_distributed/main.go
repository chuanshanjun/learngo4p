package main

import (
	"fmt"

	worker "chuanshan.github.com/learngo4p/crawler_distributed/worker/client"

	"chuanshan.github.com/learngo4p/crawler/zhenai/parser"

	"chuanshan.github.com/learngo4p/crawler/engine"
	"chuanshan.github.com/learngo4p/crawler/scheduler"
	"chuanshan.github.com/learngo4p/crawler_distributed/config"
	itemsaver "chuanshan.github.com/learngo4p/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.Concurrent{
		// 用指针...
		WorkerCount:      100,
		Scheduler:        &scheduler.QueuedScheduler{},
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://m.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
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

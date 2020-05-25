package main

import (
	"fmt"

	"chuanshan.github.com/learngo4p/crawler/engine"
	"chuanshan.github.com/learngo4p/crawler/parser"
	"chuanshan.github.com/learngo4p/crawler/scheduler"
	"chuanshan.github.com/learngo4p/crawler_distributed/config"
	"chuanshan.github.com/learngo4p/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
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

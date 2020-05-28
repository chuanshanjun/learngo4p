package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"

	"chuanshan.github.com/learngo4p/crawler_distributed/rpcsupport"

	worker "chuanshan.github.com/learngo4p/crawler_distributed/worker/client"

	"chuanshan.github.com/learngo4p/crawler/zhenai/parser"

	"chuanshan.github.com/learngo4p/crawler/engine"
	"chuanshan.github.com/learngo4p/crawler/scheduler"
	"chuanshan.github.com/learngo4p/crawler_distributed/config"
	itemsaver "chuanshan.github.com/learngo4p/crawler_distributed/persist/client"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(
		*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	// 目前worker只有一个
	processor := worker.CreateProcessor(pool)

	// 100个worker全部扔给了一个worker
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

func createClientPool(hosts []string) chan *rpc.Client {
	// 原先client是我私有的数据，不和别人共享，但我现在通过channel分发给别人
	var rpcClients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			rpcClients = append(rpcClients, client)
			log.Print("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	// 分发的话就在goroutine完成，因为收发的时候两端都需要监听
	go func() {
		for {
			// 但此时发一轮就结束了，所以外面需要再套一个for
			for _, c := range rpcClients {
				out <- c
			}
		}
	}()
	return out
}

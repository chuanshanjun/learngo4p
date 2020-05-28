package client

import (
	"fmt"

	"chuanshan.github.com/learngo4p/crawler_distributed/worker"

	"chuanshan.github.com/learngo4p/crawler/engine"
	"chuanshan.github.com/learngo4p/crawler_distributed/config"
	"chuanshan.github.com/learngo4p/crawler_distributed/rpcsupport"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(
		fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	// rpc通过函数式编程，放到了返回的函数中
	return func(req engine.Request) (engine.ParseResult, error) {
		// 1 将request序列化
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		// 2 调用
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		// 3 没有出错就反序列化,result
		return worker.DeserializeResult(sResult), nil
	}, nil
}

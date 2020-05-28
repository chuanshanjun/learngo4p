package client

import (
	"net/rpc"

	"chuanshan.github.com/learngo4p/crawler_distributed/worker"

	"chuanshan.github.com/learngo4p/crawler/engine"
	"chuanshan.github.com/learngo4p/crawler_distributed/config"
)

// clients []*rpc.Client如果使用slice的话，就会面临传统并发编程的问题
// 100个worker去抢client，那么就需要加锁，效率就低了
// Go语言的话，使用chan，clientChan chan *rpc.Client
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	// clients不用自己建，而是外面传进来
	//client, err := rpcsupport.NewClient(
	//	fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}
	// rpc通过函数式编程，放到了返回的函数中
	return func(req engine.Request) (engine.ParseResult, error) {
		// 1 将request序列化
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		// 2 调用
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		// 3 没有出错就反序列化,result
		return worker.DeserializeResult(sResult), nil
	}
}

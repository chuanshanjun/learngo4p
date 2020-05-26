package main

import (
	"fmt"
	"testing"
	"time"

	"chuanshan.github.com/learngo4p/crawler_distributed/config"

	"chuanshan.github.com/learngo4p/crawler_distributed/worker"

	"chuanshan.github.com/learngo4p/crawler_distributed/rpcsupport"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(
		host, worker.CrawlService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://m.zhenai.com/u/1115584152",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "活出高姿态",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}

package main

import (
	"testing"
	"time"

	"chuanshan.github.com/learngo4p/crawler/engine"
	"chuanshan.github.com/learngo4p/crawler/model"

	"chuanshan.github.com/learngo4p/crawler_distributed/rpcsupport"
)

// 用一个协程去开启一个Server,自己模拟客户端去请求
func TestItemSaver(t *testing.T) {
	const host = ":1234"

	//json, err := json.Marshal(item)
	//if err != nil {
	//	fmt.Println(json)
	//}

	// Start ItemSaverServer
	// 一开始报错是因为，goroutine还要去连接elastic，所以还没有起来
	// 此时client去连接，肯定失败了，
	// 现实中我们需要一个强壮的机制，当RPC服务启动后，我们告诉外面
	// 服务已经启动了，这样更加安全
	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	// Start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Url:  "http://hanxiao",
		Type: "zhenai",
		Id:   "abcDDE",
		Payload: model.Profile{
			Name:      "涵笑",
			Gender:    "女士",
			Age:       45,
			Height:    159,
			Weight:    52,
			Income:    "5-8",
			Marriage:  "离异",
			Education: "高中及以下",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
